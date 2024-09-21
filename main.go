package main

import (
	"log"
	"os"

	"github.com/vctaragao/mp3/internal/decoder"
)

func main() {
	f, err := os.Open("Daydream - Soobin Hoang SonThaoboy (Hiderway Remix).mp3")
	if err != nil {
		panic(err)
	}

	d := decoder.NewDecoder()
	mp3File, err := d.Decode(f)
	if err != nil {
		log.Fatalf("decoding mp3 file: %v", err)
	}

	mp3File.ShowHeader()
}
