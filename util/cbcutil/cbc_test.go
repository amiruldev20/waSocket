package cbcutil

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := []byte("amiruldev1337")
	plain := []byte("wellcome to waSocket")

	cipher, err := Encrypt(key, nil, plain)
	if err != nil {
		t.Fail()
	}

	p, err := Decrypt(key, nil, cipher)
	if err != nil {
		t.Fail()
	}

	if !bytes.Equal(plain, p) {
		t.Fail()
	}
}
