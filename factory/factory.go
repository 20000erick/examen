package factory

import (
	"errors"
	"examen/translate"
	"examen/utils"
)

//Factory para creacion de translate que cumplan con ITranslate
func Factory(texto string, input string, output string) (translate.ITranslate, error) {
	switch input {
	case utils.TEXT:
		switch output {
		case utils.BINARY:
			return &translate.TranslateBinaryToText{Texto: texto}, nil
		case utils.MORSE:
			return &translate.TranslateTextToMorse{Texto: texto}, nil
		default:
			return nil, errors.New(utils.INVALID_FORMAT_DESTINO)
		}
	case utils.MORSE:
		switch output {
		case utils.BINARY:
			return &translate.TranslateMorseToBinary{Texto: texto}, nil
		case utils.TEXT:
			return &translate.TranslateMorseToText{Texto: texto}, nil
		default:
			return nil, errors.New(utils.INVALID_FORMAT_DESTINO)
		}
	case utils.BINARY:
		switch output {
		case utils.MORSE:
			return &translate.TranslateBinaryToMorse{Texto: texto}, nil
		case utils.TEXT:
			return &translate.TranslateBinaryToText{Texto: texto}, nil
		default:
			return nil, errors.New(utils.INVALID_FORMAT_DESTINO)
		}
	default:
		return nil, errors.New(utils.INVALID_FORMAT_ANY)
	}

}
