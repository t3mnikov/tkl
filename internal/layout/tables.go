package layout

// Парные ряды: одна позиция = одна клавиша
var pairs = []struct{ ru, en string }{
	// нижний регистр
	{"йцукенгшщзхъ", "qwertyuiop[]"},
	{"фывапролджэ\\", "asdfghjkl;'\\"},
	{"ячсмитьбю.", "zxcvbnm,./"},

	// верхний регистр
	{"ЙЦУКЕНГШЩЗХЪ", "QWERTYUIOP{}"},
	{"ФЫВАПРОЛДЖЭ/", "ASDFGHJKL:\"|"},
	{"ЯЧСМИТЬБЮ,", "ZXCVBNM<>?"},

	// дополнительные знаки
	{"ё", "`"},
	{"Ё", "~"},
	{"1234567890-=", "1234567890-="},
	{"!\"№;%:?*()_+", "!@#$%^&*()_+"},
}

var (
	ruToEn = map[rune]rune{}
	enToRu = map[rune]rune{}
)

// init подготавливает мапы для раскладок ru и en
func init() {
	for _, p := range pairs {
		ru := []rune(p.ru)
		en := []rune(p.en)

		if len(ru) != len(en) {
			panic("tkl: mismatch pairs length: " + p.ru + " / " + p.en)
		}

		for i := range ru {
			ruToEn[en[i]] = ru[i]
			enToRu[ru[i]] = en[i]
		}
	}
}
