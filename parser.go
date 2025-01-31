package main

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
	return Info{
		files:       nil,
		length:      readIntValue(info, "length"),
		name:        readStringValue(info, "name"),
		pieceLength: readIntValue(info, "piece length"),
		pieces:      readStringValue(info, "pieces"),
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
