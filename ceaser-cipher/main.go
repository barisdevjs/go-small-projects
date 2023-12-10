package main

import (
	"fmt"
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetterFn(key int, letter string) (result string) {
	runes := []rune(letter)
	// ex : 21:26
	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	// ex : 0:21
	leftOversLetter := string(runes[0 : len(letter)-key])
	// returns new mapped alphabet
	return fmt.Sprintf(`%s%s`, lastLetterKey, leftOversLetter)
}

func encrypt(key int, plainText string) (result string) {
	hashLetter := hashLetterFn(key, alphabet)
	var hashedStr = ""
	findOne := func(r rune) rune {
		pos := strings.Index(alphabet, string([]rune{r}))
		if pos != -1 {
			hashedStr = hashedStr + string(hashLetter[pos])
			return r
		}
		return r
	}
	_ = strings.Map(findOne, plainText)
	return hashedStr
}

func decrypt(key int, encryptedText string) (result string) {
	hashLetter := hashLetterFn(key, alphabet)
	var hashedStr = ""
	findOne := func(r rune) rune {
		pos := strings.Index(hashLetter, string([]rune{r}))
		if pos != -1 {
			hashedStr = hashedStr + string(alphabet[pos])
			return r
		}
		return r
	}
	_ = strings.Map(findOne, encryptedText)
	return hashedStr
}

func main() {
	plainText := "HELLOWORLD"
	fmt.Println("plain text", plainText)
	encryptedText := encrypt(5, plainText)
	fmt.Println("encrypted text", encryptedText)
	decryptedText := decrypt(5, encryptedText)
	fmt.Println("decrypted text", decryptedText)
}
