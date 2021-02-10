package translate

type TranslateMorseToBinary struct {
	Texto string
}

func (t *TranslateMorseToBinary) TranslateInput() string {
	return "Mock Morse to Binary"
}
