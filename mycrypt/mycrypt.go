package mycrypt

import (
	"fmt"
)

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := sokIAlfabetet(melding[i], alphabet)
		if indeks+chiffer >= len(alphabet) {
			kryptertMelding[i] = alphabet[indeks+chiffer-len(alphabet)]
		} else {
			kryptertMelding[i] = alphabet[indeks+chiffer]
		}
	}
	return kryptertMelding
}

func DeKrypter(kryptertMelding []rune, alphabet []rune, chiffer int) []rune {
	melding := make([]rune, len(kryptertMelding))
	for i := 0; i < len(kryptertMelding); i++ {
		indeks := sokIAlfabetet(kryptertMelding[i], alphabet)
		if indeks-chiffer < 0 {
			melding[i] = alphabet[indeks-chiffer+len(alphabet)]
		} else {
			melding[i] = alphabet[indeks-chiffer]
		}
	}
	return melding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		if symbol == alfabet[i] {
			return i
		}
	}
	return -1
}

func main() {
	alphabet := ALF_SEM03
	chiffer := 3

	originalMessage := []rune("ping")
	fmt.Println("Original message:", string(originalMessage))

	encryptedMessage := Krypter(originalMessage, alphabet, chiffer)
	fmt.Println("Encrypted message:", string(encryptedMessage))

	decryptedMessage := DeKrypter(encryptedMessage, alphabet, chiffer)
	fmt.Println("Decrypted message:", string(decryptedMessage))
}

