package main

import (
	"fmt"
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
	mp3File.ShowID3v2Header()
	fmt.Println()
	mp3File.ShowFramesHeader()

	// header := make([]byte, 256)
	// if _, err := f.ReadAt(header, 0); err != nil {
	// 	panic(err)
	// }
	//
	// for i, b := range header {
	// 	fmt.Printf("%d: %08b - %s - %d\n", i, b, string(b), b)
	// }
	//
}
