package day2

import (
	"fmt"
	"regexp"
	"strconv"
)

type Game struct {
	id      int
	samples []Sample
}

type Sample struct {
	red   int
	green int
	blue  int
}

func ParseGame(bytes []byte) Game {
	tokens := lex(bytes)
	return parse(tokens)
}

func lex(bytes []byte) [][]byte {
	var tokens [][]byte

	for len(bytes) > 0 {
		token := lexAlphanum(bytes)
		if len(token) > 0 {
			tokens = append(tokens, token)
			bytes = bytes[len(token):]
			continue
		}

		char := bytes[0]
		switch char {
		case ':':
			fallthrough
		case ',':
			fallthrough
		case ';':
			tokens = append(tokens, []byte{char})
			bytes = bytes[1:]
		case ' ':
			bytes = bytes[1:]
		default:
			panic(fmt.Errorf("unexpected character: %v", char))
		}
	}

	return tokens
}

func lexAlphanum(bytes []byte) []byte {
	var str []byte
	r := regexp.MustCompile("[a-zA-Z0-9]")

	for _, char := range bytes {
		if !r.Match([]byte{char}) {
			break
		}
		str = append(str, char)
	}

	return str
}

func parse(tokens [][]byte) Game {
	game := Game{}

	for len(tokens) > 0 {
		token := tokens[0]
		switch string(token) {
		case "Game":
			game.id, tokens = parseGameId(tokens[1:])
		case ":":
			fallthrough
		case ";":
			var sample Sample
			sample, tokens = parseGameSample(tokens[1:])
			game.samples = append(game.samples, sample)
		default:
			panic(fmt.Errorf("unexpected token: %s", token))
		}
	}

	return game
}

func parseGameId(tokens [][]byte) (int, [][]byte) {
	id := parseInt(tokens[0])
	return id, tokens[1:]
}

func parseInt(token []byte) int {
	id, err := strconv.Atoi(string(token))
	if err != nil {
		panic(err)
	}
	return id
}

func parseGameSample(tokens [][]byte) (Sample, [][]byte) {
	var sample Sample

	for len(tokens) > 0 {
		token := tokens[0]

		color, amount := parseCube(tokens)
		if len(color) > 0 {
			switch color {
			case "red":
				sample.red = amount
			case "green":
				sample.green = amount
			case "blue":
				sample.blue = amount
			default:
				fmt.Println(fmt.Errorf("unexpected color: %s", color))
			}

			tokens = tokens[2:]
			continue
		}

		char := token[0]
		if char == ',' {
			tokens = tokens[1:]
		} else if char == ';' {
			break
		}
	}

	return sample, tokens
}

func parseCube(tokens [][]byte) (string, int) {
	var amount int
	var color string

	for len(tokens) > 0 {
		token := tokens[0]
		char := token[0]

		if char >= 48 && char <= 57 {
			amount = parseInt(token)
			tokens = tokens[1:]
		} else if char >= 97 && char <= 122 {
			color = string(token)
			tokens = tokens[1:]
		} else if char == ',' || char == ';' {
			break
		}
	}

	return color, amount
}
