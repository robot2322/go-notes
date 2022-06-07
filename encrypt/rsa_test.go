package encrypt

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRsa(t *testing.T) {

	data, _ := RsaEncrypt([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	origData, _ := RsaDecrypt(data)
	fmt.Println(string(origData))

}
