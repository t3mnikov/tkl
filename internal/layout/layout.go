package layout

import (
	"errors"
	"strings"
)

type Layout string

var ErrUnknownLayout = errors.New("неизвестная раскладка")
var ErrAbsentMap = errors.New("отсутствует карта раскладки")

const (
	RuLayout Layout = "RU"
	EnLayout Layout = "EN"
)

// Convert конвертирует текст на раскладку Layout
func Convert(target Layout, s string) (string, error) {
	sb := strings.Builder{}

	var m map[rune]rune
	switch target {
	case RuLayout:
		m = ruToEn
	case EnLayout:
		m = enToRu
	default:
		return "", ErrAbsentMap
	}

	for _, r := range s {
		if mapped, ok := m[r]; ok {
			sb.WriteRune(mapped)
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String(), nil
}

// ParseLayout парсит полученную раскладку в Layout
func ParseLayout(s string) (Layout, error) {
	switch strings.ToUpper(s) {
	case "RU":
		return RuLayout, nil
	case "EN":
		return EnLayout, nil
	default:
		return "", ErrUnknownLayout
	}
}
