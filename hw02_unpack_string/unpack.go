package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var (
		temp rune
		builder strings.Builder
	)
	runes := []rune(s)

	for i, r := range runes {
		if unicode.IsDigit(r) {
			if temp == rune(0) {
				return "", ErrInvalidString
			}

			repeatCount, _ := strconv.Atoi(string(r))

			if repeatCount != 0 {
				builder.WriteString(strings.Repeat(string(temp), repeatCount))
			}

			temp = rune(0)
		} else {
			if temp != rune(0) {
				builder.WriteRune(temp)
			}
			if i == len(runes) - 1 {
				builder.WriteRune(r)
			}
			temp = r
		}
	}

	return builder.String(), nil
}
