package main

import (
	"crypto/sha1"
	"fmt"
)

type BitTorrent struct {
	announce string
	info     Info
}

type Info struct {
	files       []File
	length      int
	name        string
	pieceLength int
	pieces      string
	hash        []byte
}

type File struct {
	length int
	path   []string
}

func ParseBitTorrentFile(content string) BitTorrent {
	result, _ := decode(content)
	return BitTorrent{
		announce: readStringValue(result, "announce"),
		info:     ParseInfoSection(result.(map[string]interface{})),
	}
}

func ParseInfoSection(content map[string]interface{}) Info {
	info := content["info"]
	hash := calculateInfoHash([]byte(fmt.Sprintf("%d%s%d%s", readIntValue(info, "length"),
		readStringValue(info, "name"), readIntValue(info, "piece length"), readStringValue(info, "pieces"))))
	return Info{
		files:       nil,
		length:      readIntValue(info, "length"),
		name:        readStringValue(info, "name"),
		pieceLength: readIntValue(info, "piece length"),
		pieces:      readStringValue(info, "pieces"),
		hash:        hash,
	}
}

func readStringValue(dict interface{}, key string) string {
	result := dict.(map[string]interface{})[key]
	return result.(string)
}

func readIntValue(dict interface{}, key string) int {
	result := dict.(map[string]interface{})[key]
	return result.(int)
}

func calculateInfoHash(input []byte) []byte {
	hasher := sha1.New()
	hasher.Write(input)
	return hasher.Sum(nil)
}
