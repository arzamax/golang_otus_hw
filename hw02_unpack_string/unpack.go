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
		temp    rune
		builder strings.Builder
	)

	for _, r := range s {
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

			temp = r
		}
	}

	if temp != rune(0) {
		builder.WriteRune(temp)
	}

	return builder.String(), nil
}
