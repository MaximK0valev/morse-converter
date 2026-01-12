package service

import (
	"errors"
	"strings"

	"github.com/MaximK0valev/morse-converter/pkg/morse"
)

// AutoConvert trims input and converts it either from Morse to text or from text to Morse.
// It returns an error if the input is empty or contains only whitespace.
func AutoConvert(data string) (string, error) {
	str := strings.TrimSpace(data)

	if str == "" {
		return "", errors.New("Входная строка пуста")
	}
	isMorse := true

	for _, s := range str {
		if !(s == '.' || s == '-' || s == ' ' || s == '/' || s == '\n') {
			isMorse = false
			break
		}
	}
	if isMorse {
		result := morse.ToText(str)
		return result, nil
	} else {
		result := morse.ToMorse(str)
		return result, nil
	}

}
