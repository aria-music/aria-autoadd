package main

type AddToPlaylistRequest struct {
	Key      string                   `json:"key"`
	OP       string                   `json:"op"`
	Postback string                   `json:"postback,omitempty"`
	Data     AddToPlaylistRequestData `json:"data,omitempty"`
}

type AddToPlaylistRequestData struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}
