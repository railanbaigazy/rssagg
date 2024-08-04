package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/railanbaigazy/rssagg/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		publishedAt, err := parsePublishedAt(item.PubDate)
		if err != nil {
			log.Printf("Couldn't parse date %v with err %v", item.PubDate, err)
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			Url:         item.Link,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("Failed to create post:", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}

func parsePublishedAt(pubDate string) (time.Time, error) {
	var publishedAt time.Time
	var err error

	dateFormats := []string{
		time.RFC1123Z,                   // e.g., "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC1123,                    // e.g., "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC822Z,                    // e.g., "02 Jan 06 15:04 -0700"
		time.RFC822,                     // e.g., "02 Jan 06 15:04 MST"
		time.RFC3339,                    // e.g., "2006-01-02T15:04:05Z07:00"
		time.RFC3339Nano,                // e.g., "2006-01-02T15:04:05.999999999Z07:00"
		"2006-01-02T15:04:05Z",          // e.g., "2006-01-02T15:04:05Z"
		"Mon, 2 Jan 2006 15:04:05 MST",  // e.g., "Mon, 2 Jan 2006 15:04:05 MST"
		"Mon, 02 Jan 2006 15:04:05 MST", // e.g., "Mon, 02 Jan 2006 15:04:05 MST"
	}

	for _, format := range dateFormats {
		publishedAt, err = time.Parse(format, pubDate)
		if err == nil {
			return publishedAt, nil
		}
	}

	return publishedAt, fmt.Errorf("unable to parse date: %s", pubDate)
}
