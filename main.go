package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	content := string(b)
	parsed := ParseBitTorrentFile(content)
	GetPeersList(strings.Split(parsed.announce, "//")[1], parsed)
}
