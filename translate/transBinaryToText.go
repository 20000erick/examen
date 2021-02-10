package translate

type TranslateBinaryToText struct {
	Texto string
}

func (t *TranslateBinaryToText) TranslateInput() string {
	return "Mock Binary to Text"
}
