package layout_test

import (
	"testing"

	"github.com/t3mnikov/tkl/internal/layout"
)

func TestToRuLayoutConvert(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected string
	}{
		{"ghbdtn", "привет"},
		{"GhjUhfvvF nTkTgtHTLfx yF pfDnhf", "ПроГраммА тЕлЕпеРЕДач нА заВтра"},
		{"jlyf;ls z bp ktce dsitk - ,sk cbkmysq vjhjp", "однажды я из лесу вышел - был сильный мороз"},
		{"ghbdtn? vbh", "привет, мир"},
		{"qwertyuiop[]", "йцукенгшщзхъ"},
		{"asdfghjkl;'\\", "фывапролджэ\\"},
		{"zxcvbnm,./", "ячсмитьбю."},
		{"`~", "ёЁ"},
		{"Gfkfnf #5", "Палата №5"},
		{"привет", "привет"},
	}

	for _, tc := range testCases {
		output, err := layout.Convert(layout.RuLayout, tc.Input)
		if err != nil {
			t.Error(err)
		}

		if output != tc.Expected {
			t.Errorf("expected %q, got %q", tc.Expected, output)
		}
	}
}

func TestToEnLayoutConvert(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected string
	}{
		{"привет", "ghbdtn"},
		{"пщщв ьщктштпб ышк", "good morning, sir"},
		{"Фдерщгпр ше**`ы ыгззщыув ещ иу Ыфкфр’ы куызщтышишдшен", "Although it**`s supposed to be Sarah’s responsibility"},
		{"Ш туув ф 200ер фззду", "I need a 200th apple"},
		{"ghbdtn", "ghbdtn"},
	}

	for _, tc := range testCases {
		output, err := layout.Convert(layout.EnLayout, tc.Input)
		if err != nil {
			t.Error(err)
		}

		if output != tc.Expected {
			t.Errorf("expected %q, got %q", tc.Expected, output)
		}
	}
}

func TestErrorConvert(t *testing.T) {
	testCases := []struct {
		layout.Layout
		Input    string
		Expected string
	}{
		{layout.Layout("ES"), "привет", "ghbdtn"},
	}

	for _, tc := range testCases {
		_, err := layout.Convert(tc.Layout, tc.Input)
		if err == nil {
			t.Error(err)
		}
	}
}

func TestParseLayout(t *testing.T) {
	layouts := []string{"ru", "Ru", "rU", "RU", "en", "En", "eN", "EN"}
	for _, l := range layouts {
		_, err := layout.ParseLayout(l)
		if err != nil {
			t.Error(err)
		}

	}
}

func TestErrorParseLayout(t *testing.T) {
	layouts := []string{"es", "ES"}
	for _, l := range layouts {
		_, err := layout.ParseLayout(l)
		if err == nil {
			t.Error(err)
		}
	}
}
