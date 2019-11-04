package main

import (
	"encoding/json"
	"log"
	"sync"

	"golang.org/x/net/websocket"
)

var playlist string

var key string
var lock sync.Mutex
var ws *websocket.Conn
var enc *json.Encoder

// NewPlayerConn establishes new player websocket connection
func NewPlayerConn(finish chan<- struct{}) {
	ws, err := websocket.Dial("wss://sarisia.cc/player/", "", "http://localhost/")
	if err != nil {
		log.Printf("Failed to establish player connection: %v\n", err)
		finish <- struct{}{}
		return
	}

	log.Println("Started!")

	dec := json.NewDecoder(ws)
	enc = json.NewEncoder(ws)
	for {
		var base BaseResponse
		err = dec.Decode(&base)
		if err != nil {
			log.Printf("Failed to parse json: %v\n", err)
			continue
		}

		go handlePacket(&base)
	}
}

func handlePacket(base *BaseResponse) {
	log.Printf("Received: %s\n", base.Type)

	switch base.Type {
	case "hello":
		key = base.Key
	case "event_player_state_change":
		var state EventPlayerStateChangeData
		err := json.Unmarshal(base.Data, &state)
		if err != nil {
			log.Printf("Failed to parse data as event_player_state_change: %v\n", err)
			return
		}

		if state.State != "playing" {
			return
		}

		sendToSocket(&AddToPlaylistRequest{
			Key:      key,
			OP:       "add_to_playlist",
			Postback: "SentFromAriaAutoAdd",
			Data: AddToPlaylistRequestData{
				Name: playlist,
				URI:  state.Entry.URI,
			},
		})
	}
}

func sendToSocket(data interface{}) {
	lock.Lock()
	defer lock.Unlock()

	err := enc.Encode(data)
	if err != nil {
		log.Printf("Failed to send to socket: %v\n", err)
	}
}

func setPlaylist(p string) {
	playlist = p
}
