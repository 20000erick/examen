package main

import (
	"examen/factory"
	"fmt"
	"os"
)

func main() {

	var texto,inputFormat,outputFormat string
	fmt.Print("Digite el texto a transformar: ")
	_, err := fmt.Scan(&texto)
	if err != nil {
		fmt.Printf("error al leer el texto: %v", err)
		os.Exit(1)
	}
	fmt.Print("Digite el inputFormat (MORSE,TEXT,BINARY): ")
	_, err = fmt.Scan(&inputFormat)
	if err != nil {
		fmt.Printf("error al leer el inputFormat: %v", err)
		os.Exit(1)
	}
	fmt.Print("Digite el outputFormat (MORSE,TEXT,BINARY): ")
	_, err = fmt.Scan(&outputFormat)
	if err != nil {
		fmt.Printf("error al leer el outputFormat: %v", err)
		os.Exit(1)
	}
//parametros validos para formato: MORSE,TEXT,BINARY
	fmt.Println(Translate(texto,inputFormat,outputFormat))
}
//Funcion que llama a una  de acuerdo a  formato origen y formato destino
func Translate (textoATraducir string, formatoOrigen string, formatoDestino string) string {
	resp, err := factory.Factory(textoATraducir, formatoOrigen, formatoDestino)
	if err != nil {
		return err.Error()
	}
	return resp.TranslateInput()
}
