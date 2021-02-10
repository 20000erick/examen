package translate

type TranslateMorseToText struct {
	Texto string
}

func (t *TranslateMorseToText) TranslateInput() string {
	return "Mock Morse to Text"
}
