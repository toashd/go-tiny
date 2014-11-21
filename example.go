package main

import (
	"fmt"
	"github.com/toashd/go-tiny/tiny"
)

func main() {
	// Url to decode (id)
	var url = 1337

	// Creates new TinyUrl
	t := tiny.New()

	// Encodes url with length of 5
	var enc = t.EncodeUrl(url, 5)

	// Prints encoded url
	fmt.Printf("Encoded url: %s\n", enc)

	// Prints decoded url
	fmt.Printf("Decoded url: %v\n", t.DecodeUrl(enc))
}

