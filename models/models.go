package models

type SearchRequest struct {
	Query string `json:"query"`
}

type ErrorResponse struct {
	Description string `json:"description,omitempty"`
}

type GetVideoResponse struct {
	Status string
	Links  Links          `json:"_links"`
	Videos []YoutubeVideo `json:"videos"`
	Error  ErrorResponse  `json:"error,omitempty"`
}

type Links struct {
	Next     string `json:"next"`
	Previous string `json:"prev,omitempty"`
	Base     string `json:"base"`
}

type SearchResponse struct {
	Status string
	Videos []YoutubeVideo `json:"videos"`
	Error  ErrorResponse  `json:"error,omitempty"`
}

type Thumbnail struct {
	Default string `json:"default"`
	Medium  string `json:"medium"`
	High    string `json:"high"`
}

type YoutubeVideo struct {
	Title        string    `json:"title"`
	ChannelId    string    `json:"channel_id"`
	Description  string    `json:"description"`
	ChannelTitle string    `json:"channel_title"`
	VideoId      string    `json:"video_id"`
	Thumbnails   Thumbnail `json:"thumbnails"`
	Kind         string    `json:"kind"`
	PublishedAt  string    `json:"publisheAt"`
}
