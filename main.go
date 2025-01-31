package main

import (
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	content := string(b)
	parsed := ParseBitTorrentFile(content)
	fmt.Println(parsed)
}
