package main

import (
  "regexp"
  "strconv"
  "fmt"
)

type BitTorrentV2 struct {
  announce string
  createdBy string
  creationDate string
}

func parseString(input string, pos int) (string, int) {
  fmt.Println(input)
  fmt.Println(pos)
  r, _ := regexp.Compile("[0-9]+:[a-zA-z]")
  length_str := r.FindString(input[pos:])
  fmt.Println(length_str)
  length, _ := strconv.Atoi(length_str[:len(length_str)-2])

  return input[pos + len(length_str) - 1:pos + len(length_str) - 1 + length], length
}

func parseInteger(input string, pos int) (int, int) {
  r, _ := regexp.Compile("i[-]?[0-9]+e")
  length_str := r.FindString(input[pos:])
  res, _ := strconv.Atoi(length_str[1:len(length_str)-1])
  return res, len(length_str[1:len(length_str)-1])
}

func parseList(input string, pos int) ([]interface{}, int) { 
  r, _ := regexp.Compile("l[0-9]+:.*e")
  length_str := r.FindString(input[pos:])
  list := input[pos+1:len(length_str)]

  var tokens []interface{}
  for i := 0; i < len(list); i++ {
    res, _ := decode(list[i:])
    tokens = append(tokens, res)
  }

  return tokens, len(list)
}

func parseDictionary(input string, pos int) (map[string]interface{}, int) {
  return nil, 0
}

func decode(input string) (interface{}, int) {
  var decoded interface{}
  var size int

  for i := 0; i < len(input); i++ {
    ch := input[i]

    switch ch {
      case 'i':
        decoded, size = parseInteger(input, i)
        i += size - 1
        break
      case 'l':
        decoded, size = parseList(input, i)
        i += size - 1
        break
      case 'd':
        decoded, size = parseDictionary(input, i)
        i += size - 1
        break
      default:
        decoded, size = parseString(input, i)
        i += size - 1
        break
    }
  }
  return decoded, 0
}
