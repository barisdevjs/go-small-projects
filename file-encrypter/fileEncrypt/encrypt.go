package fileEncrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"

	"github.com/xdg-go/pbkdf2"
)

func Encrypt(file string, password []byte) {
	// Check for the file if exist
	if _, err := os.Stat(file); os.IsNotExist(err) {
		panic(err.Error())
	}

	// Reading the file
	plainText, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// Create nonce
	key := password
	nonce := make([]byte, 12)

	// Randomize the nonce
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// dk ==> derivedKey
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	// cipher block
	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	//Galois Counter Mode
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	cipherText := aesgcm.Seal(nil, nonce, plainText, nil)

	// Append the nonce to the end of file
	cipherText = append(cipherText, nonce...)

	// create a source file for encrypted data
	dstFile, err := os.Create(file)
	if err != nil {
		panic(err.Error())
	}
	defer dstFile.Close() // remove or configure the error

	// write cipherText to the new source file
	_, err = io.Copy(dstFile, bytes.NewReader(cipherText))
	if err != nil {
		panic(err.Error())
	}

}

func Decrypt(file string, password []byte) {

	// Check for the encrypted file
	if _, err := os.Stat(file); os.IsNotExist(err) {
		panic(err.Error())
	}

	//
	ciphertext, err := os.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}

	key := password
	salt := ciphertext[len(ciphertext)-12:]
	str := hex.EncodeToString(salt)

	// reading the nonce by last 12 digits of the encrypted file
	nonce, err := hex.DecodeString(str)
	if err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plainText, err := aesgcm.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)
	if err != nil {
		panic(err.Error())
	}

	dstFile, err := os.Create(file)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(dstFile, bytes.NewReader(plainText))
	if err != nil {
		panic(err.Error())
	}

}
