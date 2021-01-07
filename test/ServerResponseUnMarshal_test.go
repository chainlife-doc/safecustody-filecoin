package test

import (
	"fmt"
	"safecustody_filecoin"
	"testing"
)

func Test_ServerResponseUnMarshal(t *testing.T) {

	s := "{\"error\":\"\",\"data\":\"奥术大师大\"}"

	h, err := safecustody_filecoin.ServerResponseUnMarshal([]byte(s))
	fmt.Println(err)
	fmt.Println(h)
}
