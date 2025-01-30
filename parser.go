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

func parseBitTorrentFile() BitTorrent {

}

func parseFileSection() File {

}

func parseInfoSection() Info {

}
