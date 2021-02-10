package factory

import (
	"examen/translate"
	"reflect"
	"testing"
)

func TestFactory(t *testing.T) {
	type args struct {
		texto  string
		input  string
		output string
	}
	tests := []struct {
		name    string
		args    args
		want    *translate.TranslateMorseToText
		wantErr bool
	}{
		{name: "test factory ", args: args{
			texto:  "",
			input:  "MORSE",
			output: "TEXT",
		}, want:&translate.TranslateMorseToText{Texto: ""}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factory(tt.args.texto, tt.args.input, tt.args.output)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory() got = %v, want %v", got, tt.want)
			}
		})
	}
}