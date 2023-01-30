# Project Goal

To make an API to fetch latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

# Basic Requirements:

- [x] Server should call the YouTube API continuously in background (async) with some interval (say 10 seconds) for fetching the latest videos for a predefined search query and should store the data of videos (specifically these fields - Video title, description, publishing datetime, thumbnails URLs and any other fields you require) in a database with proper indexes.
- [x] A GET API which returns the stored video data in a paginated response sorted in descending order of published datetime.
- [x] A basic search API to search the stored videos using their title and description.
- [x] Dockerize the project.
- [x] It should be scalable and optimised.

# Bonus Points:

- [ ] Add support for supplying multiple API keys so that if quota is exhausted on one, it automatically uses the next available key.
- [ ] Make a dashboard to view the stored videos with filters and sorting options (optional)
- [x] Optimise search api, so that it's able to search videos containing partial match for the search query in either video title or description.
    - Ex 1: A video with title *`How to make tea?`* should match for the search query `tea how`
    
## Tech Stack
 - Golang (Iris framework)
- MongoDB
- Docker
    
## Development:
- Clone the project
```
https://github.com/themechanicalcoder/fampay-backend-assignment.git 
cd fampay-backend-assignment
```

- Add query, apikey and MongoURI in the and `dev.yaml` file
```
Name: fampay-backend-assignment
YoutubeConfig:
  ApiKey: ""
  RelevanceLanguage: "en"
  Query: "cricket"
  MaxResults: 10
WorkerConfig:
  QueryInterval: 100 
Server:
  Addr: "0.0.0.0"
  Port: 3000
DBConfig: 
  Database: "fampay"
  Collection: "video"
  MongoURI: "mongodb://mongo:27017"
  ```
  
  
- Run in development mode by entering the command
```go run main.go```

## Running with Docker Compose
When using Docker Compose,

 - Set the ApiKey to your apiKey and MongoURI variable in your dev.yaml file to
``` MONGODB_URI = mongodb://mongo:27017```
- Run:
```
docker compose build
docker compose up
```
- The app is hosted on ```http://localhost:3000```

## APIs
#### search `/v1/search`

###  sample request
```
curl --location --request GET 'localhost:3000/v1/search' \
--header 'Content-Type: application/json' \
--data-raw '{
    "query" : "cricket"
}'

```
### sample response
```
    {
    "Status": "SUCCESS",
    "videos": [
        {
            "title": "🔴LIVE CRICKET MATCH TODAY | | CRICKET LIVE | 3rd ODI | IND vs NZ LIVE MATCH TODAY | Cricket 22",
            "channel_id": "UCywHE3e3I3hDcNg1fHhXZYQ",
            "description": "LIVE CRICKET MATCH TODAY | | CRICKET LIVE | 3rd ODI | IND vs NZ LIVE MATCH TODAY | Cricket 22 LIVE CRICKET ...",
            "channel_title": "Cricketora",
            "video_id": "vqaV48YXkg4",
            "thumbnails": {
                "default": "https://i.ytimg.com/vi/vqaV48YXkg4/default.jpg",
                "medium": "https://i.ytimg.com/vi/vqaV48YXkg4/mqdefault.jpg",
                "high": "https://i.ytimg.com/vi/vqaV48YXkg4/hqdefault.jpg"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2023-01-23T18:40:53Z"
        },
        {
            "title": "🔴LIVE CRICKET MATCH TODAY | | CRICKET LIVE | 3rd ODI | IND vs NZ LIVE MATCH TODAY | Cricket 22",
            "channel_id": "UCywHE3e3I3hDcNg1fHhXZYQ",
            "description": "LIVE CRICKET MATCH TODAY | | CRICKET LIVE | 3rd ODI | IND vs NZ LIVE MATCH TODAY | Cricket 22 LIVE CRICKET ...",
            "channel_title": "Cricketora",
            "video_id": "vqaV48YXkg4",
            "thumbnails": {
                "default": "https://i.ytimg.com/vi/vqaV48YXkg4/default.jpg",
                "medium": "https://i.ytimg.com/vi/vqaV48YXkg4/mqdefault.jpg",
                "high": "https://i.ytimg.com/vi/vqaV48YXkg4/hqdefault.jpg"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2023-01-23T18:40:53Z"
        },
        {
            "title": "CRICKET INDIA",
            "channel_id": "UCcgjrzSXOBY8gP6i-tYyGKw",
            "description": "Cricket India is a Youtube Channel Where We Will Provide Latest & Genuine News Regarding Cricket. ✓ Channel Hosted ...",
            "channel_title": "CRICKET INDIA",
            "video_id": "",
            "thumbnails": {
                "default": "https://yt3.ggpht.com/ytc/AL5GRJVUyQemoi7GPOzP12qwmzsBAwImYoL77PSto1kdiw=s88-c-k-c0xffffffff-no-rj-mo",
                "medium": "https://yt3.ggpht.com/ytc/AL5GRJVUyQemoi7GPOzP12qwmzsBAwImYoL77PSto1kdiw=s240-c-k-c0xffffffff-no-rj-mo",
                "high": "https://yt3.ggpht.com/ytc/AL5GRJVUyQemoi7GPOzP12qwmzsBAwImYoL77PSto1kdiw=s800-c-k-c0xffffffff-no-rj-mo"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2019-09-14T11:16:36Z"
        },
        {
            "title": "CRICKET INDIA",
            "channel_id": "UCcgjrzSXOBY8gP6i-tYyGKw",
            "description": "Cricket India is a Youtube Channel Where We Will Provide Latest & Genuine News Regarding Cricket. ✓ Channel Hosted ...",
            "channel_title": "CRICKET INDIA",
            "video_id": "",
            "thumbnails": {
                "default": "https://yt3.ggpht.com/ytc/AL5GRJVUyQemoi7GPOzP12qwmzsBAwImYoL77PSto1kdiw=s88-c-k-c0xffffffff-no-rj-mo",
                "medium": "https://yt3.ggpht.com/ytc/AL5GRJVUyQemoi7GPOzP12qwmzsBAwImYoL77PSto1kdiw=s240-c-k-c0xffffffff-no-rj-mo",
                "high": "https://yt3.ggpht.com/ytc/AL5GRJVUyQemoi7GPOzP12qwmzsBAwImYoL77PSto1kdiw=s800-c-k-c0xffffffff-no-rj-mo"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2019-09-14T11:16:36Z"
        },
        {
            "title": "IPL 2023 - Group A &amp; B , RCB , CSK , Start Date | Cricket Fatafat | EP 864 | MY Cricket Production",
            "channel_id": "UC07CHp5ikd-AyafR4J2LWkA",
            "description": "Watch the Episode 864 of IPL Ki Baat and Cricket Fatafat with 10 Big News of the Day including IPL Schedule and Start Date ...",
            "channel_title": "MY Cricket Production",
            "video_id": "mDws40-2Rzg",
            "thumbnails": {
                "default": "https://i.ytimg.com/vi/mDws40-2Rzg/default.jpg",
                "medium": "https://i.ytimg.com/vi/mDws40-2Rzg/mqdefault.jpg",
                "high": "https://i.ytimg.com/vi/mDws40-2Rzg/hqdefault.jpg"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2023-01-24T05:15:02Z"
        },
    ],
    "error": {}
}

```

#### videos `/v1/videos`
sample request
```
curl --location --request GET 'localhost:3000/v1/videos?limit=5&offset=5'
```
sample response

```
{
    "Status": "SUCCESS",
    "_links": {
        "next": "/v1/videos?limit=5&offset=10",
        "prev": "/v1/videos?limit=5&offset=0",
        "base": "0.0.0.0:3000"
    },
    "videos": [
        {
            "title": "Cricbuzz Live: #India clinch a thriller, beat #NewZealand, level series 1-1!",
            "channel_id": "UCSRQXk5yErn4e14vN76upOw",
            "description": "India clinch a thriller, beat New Zealand in the 2nd T20I to level the series 1-1! Dinesh Karthik, Harsha Bhogle & Gaurav Kapur ...",
            "channel_title": "Cricbuzz",
            "video_id": "ROoUOMF-sxw",
            "thumbnails": {
                "default": "https://i.ytimg.com/vi/ROoUOMF-sxw/default.jpg",
                "medium": "https://i.ytimg.com/vi/ROoUOMF-sxw/mqdefault.jpg",
                "high": "https://i.ytimg.com/vi/ROoUOMF-sxw/hqdefault.jpg"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2023-01-29T17:41:35Z"
        },
        {
            "title": "India vs New Zealand 2nd T20 Highlights: Cricket Match 2nd T20 Full Highlights | Today Match Highlig",
            "channel_id": "UCwCiYZeOc2GszCqgAQfBB2Q",
            "description": "Ind Vs NZ 2nd T20 Highlights 2023: Ind vs NZ 2nd T20 Highlights | Today Match Highlights #highlights #todaymatchhighlights ...",
            "channel_title": "NN Sports",
            "video_id": "jNcp_Vf_iS0",
            "thumbnails": {
                "default": "https://i.ytimg.com/vi/jNcp_Vf_iS0/default.jpg",
                "medium": "https://i.ytimg.com/vi/jNcp_Vf_iS0/mqdefault.jpg",
                "high": "https://i.ytimg.com/vi/jNcp_Vf_iS0/hqdefault.jpg"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2023-01-29T17:21:21Z"
        },
        {
            "title": "Ind Vs NZ 2NDT20 : रोमांचक मैच में जीता ही गया हिंदुस्तान | Hardik | Suryakumar | Arshdeep | Shubman",
            "channel_id": "UCPckg9pijh0KjJm4X0Xhviw",
            "description": "news24sports Ind Vs NZ 2NDT20 : रोमांचक मैच में जीता ही गया हिंदुस्तान | Hardik ...",
            "channel_title": "News24 Sports",
            "video_id": "kWPMBSAqRbo",
            "thumbnails": {
                "default": "https://i.ytimg.com/vi/kWPMBSAqRbo/default.jpg",
                "medium": "https://i.ytimg.com/vi/kWPMBSAqRbo/mqdefault.jpg",
                "high": "https://i.ytimg.com/vi/kWPMBSAqRbo/hqdefault.jpg"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2023-01-29T17:05:56Z"
        },
        {
            "title": "India U19 Women Lifts T20 World Cup Trophy, NZ 99/8 India RATTLED NZ, India vs New Zealand 2nd T20",
            "channel_id": "UCK4SlQg9FN9gejBpq3pNLGg",
            "description": "0:00 Intro 0:30 India Young Girls creates History 12:10 Ind vs NZ 2nd T20I Lucknow India U19 Women's Lifts T20 World Cup ...",
            "channel_title": "Sawera Pasha",
            "video_id": "JN9f8k18odo",
            "thumbnails": {
                "default": "https://i.ytimg.com/vi/JN9f8k18odo/default.jpg",
                "medium": "https://i.ytimg.com/vi/JN9f8k18odo/mqdefault.jpg",
                "high": "https://i.ytimg.com/vi/JN9f8k18odo/hqdefault.jpg"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2023-01-29T15:46:52Z"
        },
        {
            "title": "देखिए, दूसरे मैच के लिए Rohit Kohli जडेजा कि हुई भारतीय टीम में वापसी, Pandya बाहर",
            "channel_id": "UCciVT51sxyKZiHTzEPfan_g",
            "description": "indvsnz #indiavsnewzealand #highlights #fullhighlights #hardikpandya #washingtonsundar #suryakumaryadav देखिए, ...",
            "channel_title": "CRICKET UPDATES",
            "video_id": "bYvXIF6pu1c",
            "thumbnails": {
                "default": "https://i.ytimg.com/vi/bYvXIF6pu1c/default.jpg",
                "medium": "https://i.ytimg.com/vi/bYvXIF6pu1c/mqdefault.jpg",
                "high": "https://i.ytimg.com/vi/bYvXIF6pu1c/hqdefault.jpg"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2023-01-29T06:52:20Z"
        }
    ],
    "error": {}
}
```
