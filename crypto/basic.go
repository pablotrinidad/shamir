package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// Encrypt returns the result of performing AES 256 encryption using the given input string.
// A 256 bits key is generated by computing the SHA-256 function against the given string.
// The result includes the generated key and the encrypted content in that order.
func Encrypt(key string, content []byte) ([]byte, []byte, error) {
	h := sha256.New()
	h.Write([]byte(key))
	k := h.Sum(nil)

	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	cipherContent := aesGCM.Seal(nonce, nonce, content, nil)
	return k, cipherContent, nil
}
