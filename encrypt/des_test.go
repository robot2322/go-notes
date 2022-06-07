package encrypt

import (
	"fmt"
	"log"
	"testing"
)

func TestDes(t *testing.T) {

	key := []byte("2fa6c1e9")
	str := "hello world!"
	strEncrypted, err := Encrypt(str, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted:", strEncrypted)
	strDecrypted, err := Decrypt(strEncrypted, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted:", strDecrypted)

}
