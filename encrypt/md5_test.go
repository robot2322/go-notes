package encrypt

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {

	strTest := "hello world"
	fmt.Println(Encode(strTest))
	strEncrypted := "5eb63bbbe01eeed093cb22bb8f5acdc3"
	fmt.Println(Check(strTest, strEncrypted))

}
