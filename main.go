package main

import (
	"os"

	"github.com/vctaragao/mp3/internal/decoder"
)

func main() {
	f, err := os.Open("Daydream - Soobin Hoang SonThaoboy (Hiderway Remix).mp3")
	if err != nil {
		panic(err)
	}

	d := decoder.NewDecoder()
	d.Decode(f)
}
