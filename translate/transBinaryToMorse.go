package translate

type TranslateBinaryToMorse struct {
	Texto string
}

func (t *TranslateBinaryToMorse) TranslateInput() string {
	return "Mock Binary to Morse"
}
