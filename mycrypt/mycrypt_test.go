package mycrypt

import (
	"reflect"
	"testing"
)

// Testene forutsetter et alfabet, som er definert i mycrypt.go som ALF_SEM03

func TestKrypter(t *testing.T) {

	type test struct {
		inputMessage []rune
                chiffer int
		want  []rune
	}
	tests := []test{
		{inputMessage: []rune("w"), chiffer: 4, want: []rune("æ")},
		{inputMessage: []rune("0"), chiffer: 4, want: []rune("4")},
	}

	for _, tc := range tests {
		got := Krypter(tc.inputMessage, ALF_SEM03, tc.chiffer)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Feil, for chiffer %d, fikk %q, ønsket %q.", tc.chiffer, got, tc.want)
		}
	}
}

// Posisjonene i alfabetet begynner på 0 fra venstre og teller oppover mot høyre
func TestSokIAlfabetet(t *testing.T) {
	type test struct {
		input rune
		want  int
	}
	tests := []test{
		{input: 'æ', want: 26},
		{input: 'a', want: 0},
	}

	for _, tc := range tests {
		got := sokIAlfabetet(tc.input, ALF_SEM03)
		if got != tc.want {
			t.Errorf("Feil, fikk %d, ønsket %d.", got, tc.want)
		}
	}
}
