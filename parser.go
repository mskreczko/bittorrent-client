package main

import (
  "regexp"
  "strconv"
)

type BitTorrentV2 struct {
  announce string
  createdBy string
  creationDate string
}

func parseString(input string, pos int) (string, int) {
  r, _ := regexp.Compile("[0-9]+:[a-zA-z]")
  length_str := r.FindString(input[pos:])
  length, _ := strconv.Atoi(length_str[:len(length_str)-2])

  return input[pos + len(length_str) - 1:pos + len(length_str) - 1 + length], length
}

func parseInteger(input string, pos int) (int, int) {
  r, _ := regexp.Compile("i[-]?[0-9]+e")
  length_str := r.FindString(input[pos:])
  res, _ := strconv.Atoi(length_str[1:len(length_str)-1])
  return res, len(length_str[1:len(length_str)-1])
}

func parseList(input string, pos int) {
}

func parseDictionary(input string, pos int) {
  
}

func tokenize(input string) []string {
  // var tokens []string

  for i := 0; i < len(input); i++ {
    ch := input[i]

    switch ch {
      case 's':
        parseString(input, i)
      case 'i':
        parseInteger(input, i)
      case 'l':
        parseList(input, i)
      case 'd':
        parseDictionary(input, i)
    }
  }

  return []string{}
}
