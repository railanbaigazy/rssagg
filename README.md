# RSS Feed Aggregator

A Golang RESTful API Web Server for aggregating RSS feeds, powered by a PostgreSQL database.

## Getting Started

### 1. Set Up Your Environment

Create a `.env` file with the following variables:
```
PORT=your_port
DB_URL=postgres://username:password@host:port/database
```

### 2. Run Migrations

Install Goose and run the migrations:
```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir ./migrations postgres $DB_URL up
```

### 3. Start the Server

Build and start the server:
```sh
go build && ./rssagg
```

To stop the server, use `⌃C` on MacOS or `Ctrl+C` on Linux/Windows.

## API Documentation

### Health Check

**GET `/v1/healthz`**

Response:
```json
{
  "status": "ok"
}
```

### Error Simulation

**GET `/v1/err`**

Response:
```json
{
  "error": "Internal Server Error"
}
```

### Create an Account

**POST `/v1/accounts`**

Request:
```json
{
  "name": "Railan"
}
```

Response:
```json
{
  "id": "8916a240-8d65-4051-bd65-0208780f5c95",
  "created_at": "2024-08-04T14:18:30.107648+05:00",
  "updated_at": "2024-08-04T14:18:30.107649+05:00",
  "name": "Railan",
  "api_key": "3ca78b4ae9c413ac697a487787738652b8f36ded169d184a171074c9f5e57aec"
}
```

### Retrieve Account Information

**GET `/v1/accounts`**

Header:
```
Authorization: ApiKey {api_key}
```

Response:
```json
{
  "id": "8916a240-8d65-4051-bd65-0208780f5c95",
  "created_at": "2024-08-04T14:18:30.107648+05:00",
  "updated_at": "2024-08-04T14:18:30.107649+05:00",
  "name": "Railan",
  "api_key": "3ca78b4ae9c413ac697a487787738652b8f36ded169d184a171074c9f5e57aec"
}
```

### Create a Feed

**POST `/v1/feeds`**

Header:
```
Authorization: ApiKey {api_key}
```

Request:
```json
{
  "name": "The Yapping Blog",
  "url": "https://yappingblog.com/index.xml"
}
```

Response:
```json
{
  "id": "f3fd2640-19cc-4af4-88fc-8f544907bfe9",
  "created_at": "2024-08-04T14:27:54.580067+05:00",
  "updated_at": "2024-08-04T14:27:54.580067+05:00",
  "name": "The Yapping Blog",
  "url": "https://yappingblog.com/index.xml",
  "account_id": "8916a240-8d65-4051-bd65-0208780f5c95"
}
```

### Retrieve All Feeds

**GET `/v1/feeds`**

Response:
```json
[
  {
    "id": "f77a2d09-4d93-4f86-8efd-0ab52167e543",
    "created_at": "2024-08-03T17:01:49.08383+05:00",
    "updated_at": "2024-08-04T14:30:04.084801+05:00",
    "name": "The Incredible Podcast",
    "url": "https://incrediblepodcast.kz/index.xml",
    "account_id": "b753f031-5f72-4066-9411-e705de82e889"
  },
  {
    "id": "f3fd2640-19cc-4af4-88fc-8f544907bfe9",
    "created_at": "2024-08-04T14:27:54.580067+05:00",
    "updated_at": "2024-08-04T14:30:04.098321+05:00",
    "name": "The Yapping Blog",
    "url": "https://yappingblog.com/index.xml",
    "account_id": "8916a240-8d65-4051-bd65-0208780f5c95"
  }
]
```

### Follow a Feed

**POST `/v1/feed_follows`**

Header:
```
Authorization: ApiKey {api_key}
```

Request:
```json
{
  "feed_id": "{feed_id}"
}
```

Response:
```json
{
  "id": "0383a592-7cf0-455f-981a-47b0d39a0e7a",
  "created_at": "2024-08-04T14:38:05.027996+05:00",
  "updated_at": "2024-08-04T14:38:05.027996+05:00",
  "account_id": "8916a240-8d65-4051-bd65-0208780f5c95",
  "feed_id": "f3fd2640-19cc-4af4-88fc-8f544907bfe9"
}
```

### Unfollow a Feed

**DELETE `/v1/feed_follows/{feedFollowID}`**

Header:
```
Authorization: ApiKey {api_key}
```

Response:
*Status code 200*

### Retrieve All Followed Feeds

**GET `/v1/feed_follows`**

Header:
```
Authorization: ApiKey {api_key}
```

Response:
```json
[
  {
    "id": "8445b824-6d56-4af3-bec4-ebda2453f05d",
    "created_at": "2024-08-04T01:09:31.016413+05:00",
    "updated_at": "2024-08-04T01:09:31.016413+05:00",
    "account_id": "dd17e3ff-06d9-449f-8325-56ea0c8d2e0e",
    "feed_id": "f77a2d09-4d93-4f86-8efd-0ab52167e543"
  },
  {
    "id": "45d882b0-0c1a-4127-8693-fb0df107c2e0",
    "created_at": "2024-08-04T01:09:42.736072+05:00",
    "updated_at": "2024-08-04T01:09:42.736072+05:00",
    "account_id": "dd17e3ff-06d9-449f-8325-56ea0c8d2e0e",
    "feed_id": "565443a9-4eb6-4f15-beb0-400d8bf64c57"
  }
]
```

### Retrieve All Posts

**GET `/v1/posts`**

Header:
```
Authorization: ApiKey {api_key}
```

Response:
```json
[
  {
    "id": "b1173b7f-8811-42e0-b02a-0a69e9d56e85",
    "created_at": "2024-08-04T00:43:46.617504+05:00",
    "updated_at": "2024-08-04T00:43:46.617504+05:00",
    "title": "The Revolution. August 2024",
    "description": "Robots are getting over the planet",
    "url": "https://yapping.com/news/revolution-vr-2024-08/",
    "published_at": "2024-07-26T05:00:00+05:00",
    "feed_id": "f77a2d09-4d93-4f86-8efd-0ab52167e543"
  },
  {
    "id": "c8037929-eb16-4bc9-afe2-a42f65243de8",
    "created_at": "2024-08-04T00:43:46.61333+05:00",
    "updated_at": "2024-08-04T00:43:46.61333+05:00",
    "title": "Kimberly Oogway on Another Podcast?",
    "description": "Productivity guru shares the secret of being super active",
    "url": "https://anotherpodcast.com/entertain/245/",
    "published_at": "2024-07-26T05:00:00+05:00",
    "feed_id": "565443

a9-4eb6-4f15-beb0-400d8bf64c57"
  }
]
```