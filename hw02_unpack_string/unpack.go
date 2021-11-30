package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var output string
	var previousRune rune
	previousIsDigit := true
	inputLen := len(input)

	for i, r := range input {
		if unicode.IsDigit(r) {
			if previousIsDigit {
				return "", ErrInvalidString
			}
			digit, _ := strconv.Atoi(string(r))
			output += strings.Repeat(string(previousRune), digit)
		} else {
			if !previousIsDigit {
				output += string(previousRune)
			}
			if inputLen == i+1 {
				output += string(r)
			}
		}

		previousIsDigit = unicode.IsDigit(r)
		previousRune = r
	}
	return output, nil
}
