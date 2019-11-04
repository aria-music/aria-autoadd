package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("No playlist name is given. Exit.")
	}

	setPlaylist(os.Args[1])

	sigFinish := make(chan struct{})
	go NewPlayerConn(sigFinish)
	<-sigFinish
}
