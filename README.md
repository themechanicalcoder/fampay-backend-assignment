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
    
## Development:
- Clone the project
```https://github.com/themechanicalcoder/fampay-backend-assignment.git 
   cd fampay-backend-assignment
```

- Add apikey and mongodburi in the `dev.yaml` file
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
  Database: "fampay1"
  Collection: "video1"
  MongoURI: "" 
  ```
  
  
- Run in development mode
```go run main.go```


## Running with Docker Compose
When using Docker Compose,

 - Set the MONGODB_URI environment variable in your .env file to
``` MONGODB_URI = mongodb://mongo:27017```
- Run:
```docker compose build
   docker compose up
```
- Navigate to ```http://localhost:3000``` to see the app live
