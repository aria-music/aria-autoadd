package main

import "encoding/json"

type BaseResponse struct {
	Type     string
	Postback string
	Key      string
	Data     json.RawMessage
}

type EventPlayerStateChangeData struct {
	State string
	Entry struct {
		Source         string
		Title          string
		URI            string
		Thumbnail      string
		ThumbnailSmall string
		Entry          interface{}
		IsLiked        bool
		Duration       float32
		Position       float32
	}
}
