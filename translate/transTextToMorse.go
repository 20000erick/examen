package translate

type TranslateTextToMorse struct {
	Texto string
}

func (t *TranslateTextToMorse) TranslateInput() string {
	return "Mock text to morse"
}
