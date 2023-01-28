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
 - Golang 
- MongoDB
    
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
curl --location --request GET 'localhost:3000/v1/videos?limit=5&offset=3'
```
sample response

```
{
    "Status": "SUCCESS",
    "videos": [
        {
            "title": "Cricket Team India | Rohit Sharma ने ठोका 3 साल बाद ODI शतक, Suryakumar Yadav फिर flop | IND vs NZ",
            "channel_id": "UC5ebo42ydvAayGn2Z4Lf9XA",
            "description": "Cricket Team India | Rohit Sharma ने ठोका 3 साल बाद ODI शतक, Suryakumar Yadav फिर flop | IND vs NZ #indvsnz ...",
            "channel_title": "Uncut",
            "video_id": "hFGUZWMyqO8",
            "thumbnails": {
                "default": "https://i.ytimg.com/vi/hFGUZWMyqO8/default.jpg",
                "medium": "https://i.ytimg.com/vi/hFGUZWMyqO8/mqdefault.jpg",
                "high": "https://i.ytimg.com/vi/hFGUZWMyqO8/hqdefault.jpg"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2023-01-24T16:30:08Z"
        },
        {
            "title": "Pooran-Pollard Partnership | MI Emirates V/S Desert Vipers - 20 20 Cricket | 24-01-2023 | ILT20-2023",
            "channel_id": "UCyh4IJlxulhbRaZcbZVXZOA",
            "description": "Watch Live Cricket ILT20 : https://zee5.onelink.me/RlQq/840ozuv4 Subscribe to ILT20 here : https://rb.gy/03tbrt Cricket Match 20 ...",
            "channel_title": "ILT20 On Zee",
            "video_id": "uhC2RoW2iJQ",
            "thumbnails": {
                "default": "https://i.ytimg.com/vi/uhC2RoW2iJQ/default.jpg",
                "medium": "https://i.ytimg.com/vi/uhC2RoW2iJQ/mqdefault.jpg",
                "high": "https://i.ytimg.com/vi/uhC2RoW2iJQ/hqdefault.jpg"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2023-01-24T16:29:40Z"
        },
        {
            "title": "Ind vs Nz 3rd ODI: Umran Shardul ने साँस रोकने वाली रफ़तार से उड़ाए NZ के पर्क्छे जीता हारा हुआ मैच",
            "channel_id": "UC2h1GsM_Ls1Cg6M-Mk42UYQ",
            "description": "HIGHLIGHTS #LIVE #IndiavsEngland #IndvsEng #indvsengt202022 #Ipl2022 #RohitSharma​ #ViratKohli​ #JaspritBumrah​ ...",
            "channel_title": "Sports Edge Cricket",
            "video_id": "9ZqzLqddBtc",
            "thumbnails": {
                "default": "https://i.ytimg.com/vi/9ZqzLqddBtc/default.jpg",
                "medium": "https://i.ytimg.com/vi/9ZqzLqddBtc/mqdefault.jpg",
                "high": "https://i.ytimg.com/vi/9ZqzLqddBtc/hqdefault.jpg"
            },
            "kind": "youtube#searchResult",
            "publisheAt": "2023-01-24T16:05:04Z"
        },
    ],
    "error": {}
}
```
