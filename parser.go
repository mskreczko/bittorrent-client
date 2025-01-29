package main

import (
	"regexp"
	"strconv"
	"strings"
)

func parseString(input string) (string, int) {
	parts := strings.Split(input, ":")
	strLength, _ := strconv.Atoi(parts[0])

	return input[len(parts)+1 : len(parts)+strLength+1], strLength + len(parts) + 1
}

func parseInteger(input string, pos int) (int, int) {
	r, _ := regexp.Compile("i-?[0-9]+e")
	lengthStr := r.FindString(input[pos:])
	res, _ := strconv.Atoi(lengthStr[1 : len(lengthStr)-1])
	return res, len(lengthStr)
}

func parseList(input string, pos int) ([]interface{}, int) {
	r, _ := regexp.Compile("l[0-9]+:.*e")
	lengthStr := r.FindString(input[pos:])
	list := input[pos+1 : len(lengthStr)]

	var tokens []interface{}
	for i := 0; i < len(list); i++ {
		res, idx := decode(list[i:])
		tokens = append(tokens, res)
		i += idx - 1
	}

	return tokens, len(list)
}

func parseDictionary(input string, pos int) (map[string]interface{}, int) {
	return nil, 0
}

func decode(input string) (interface{}, int) {
	var decoded interface{}
	var size int

	var retIdx int
	for i := 0; i < len(input); i++ {
		ch := string([]rune(input)[i])

		switch ch {
		case "i":
			decoded, size = parseInteger(input, i)
			retIdx += size
			i += size
			break
		case "l":
			decoded, size = parseList(input, i)
			retIdx += size
			i += size
			break
		case "d":
			decoded, size = parseDictionary(input, i)
			retIdx += size
			i += size
			break
		default:
			decoded, size = parseString(input)
			retIdx += size
			i += size
			break
		}
	}
	return decoded, retIdx
}
