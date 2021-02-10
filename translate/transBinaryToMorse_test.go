package translate

import "testing"

func TestTranslateBinaryToMorse_TranslateInput(t1 *testing.T) {
	type fields struct {
		Texto string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "test binary to morse", fields: fields{Texto: ""}, want:"Mock Binary to Morse" },
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TranslateBinaryToMorse{
				Texto: tt.fields.Texto,
			}
			if got := t.TranslateInput(); got != tt.want {
				t1.Errorf("TranslateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}