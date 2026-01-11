package service

import (
	"errors"
	"strings"

	"github.com/SaidHasan-go/morse-converter/pkg/morse"
)

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
