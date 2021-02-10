package translate

type TranslateTextToBinary struct {
	Texto string
}

func (t *TranslateTextToBinary) TranslateInput() string {
	return "Mock text to Binary"
}
