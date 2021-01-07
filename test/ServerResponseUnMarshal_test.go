package test

import (
	"safecustody_filecoin"
	"testing"
)

func Test_ServerResponseUnMarshal(t *testing.T) {

	_ = safecustody_filecoin.ServerResponseUnMarshal(safecustody_filecoin.ServerResponseMarshal("", "sadasd"))

}
