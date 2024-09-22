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
	mp3File.ShowID3v2Tag()

	// header := make([]byte, 256)
	//
	// if _, err := io.ReadFull(f, header); err != nil {
	// 	panic(err)
	// }
	//
	// for i, b := range header {
	// 	fmt.Printf("%d: %08b - %s - %d\n", i, b, string(b), b)
	// }
	//
}
