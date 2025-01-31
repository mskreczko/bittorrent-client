package main

import (
	"regexp"
	"strconv"
	"strings"
)

func parseString(input string) (string, int) {
	parts := strings.Split(input, ":")
	strLength, _ := strconv.Atoi(parts[0])

	return input[len(parts[0])+1 : len(parts[0])+strLength+1], strLength + len(parts[0]) + 1
}

func parseInteger(input string) (int, int) {
	r, _ := regexp.Compile("i-?[0-9]+e")
	lengthStr := r.FindString(input)
	res, _ := strconv.Atoi(lengthStr[1 : len(lengthStr)-1])
	return res, len(lengthStr)
}

func parseList(input string) ([]interface{}, int) {
	r, _ := regexp.Compile("l[0-9]+:.*e")
	lengthStr := r.FindString(input)
	list := input[:len(lengthStr)]

	var tokens []interface{}
	i := 1
	for i < len(input) {
		if input[i] == 'e' {
			break
		}
		res, idx := decode(list[i:])
		tokens = append(tokens, res)
		i += idx
	}

	return tokens, len(list)
}

func parseDictionary(input string) (map[string]interface{}, int) {
	r, _ := regexp.Compile("d.+e")
	lengthStr := r.FindString(input)

	result := make(map[string]interface{})
	i := 1
	for i < len(input) {
		if input[i] == 'e' {
			break
		}
		key, idx := decode(input[i:])
		val, idx2 := decode(input[i+idx:])
		if key == "" || val == "" {
			break
		}
		result[key.(string)] = val
		i += idx + idx2
	}

	return result, len(lengthStr)
}

func decode(input string) (interface{}, int) {
	var decoded interface{}

	i := 0
	for i < len(input) {
		ch := string([]rune(input)[i])

		switch ch {
		case "i":
			return parseInteger(input[i:])
		case "l":
			return parseList(input[i:])
		case "d":
			return parseDictionary(input[i:])
		default:
			return parseString(input[i:])
		}
	}
	return decoded, i
}
