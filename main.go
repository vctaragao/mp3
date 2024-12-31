package main

import (
	"log"
	"os"

	"github.com/vctaragao/mp3/internal/mp3"
)

func main() {
	f, err := os.Open("Daydream - Soobin Hoang SonThaoboy (Hiderway Remix).mp3")
	if err != nil {
		panic(err)
	}

	d := mp3.NewDecoder()

	_, err = d.Decode(f)
	if err != nil {
		log.Fatalf("decoding mp3 file: %v", err)
	}
}
