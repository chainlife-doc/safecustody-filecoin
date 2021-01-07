package safecustody_filecoin

import (
	"encoding/json"
	"errors"
	"github.com/filecoin-project/go-address"
	"time"
)

func MarshalBroadCastData(hex []byte) []byte {
	p := Hex{
		Hex: string(hex),
		T:   time.Now().Unix(),
	}

	b, _ := json.Marshal(p)

	return b
}

func MarshalGetNonceData(addr address.Address) []byte {
	p := Nonce{FromAddr: addr, T: time.Now().Unix()}

	b, _ := json.Marshal(p)

	return b
}

func MarshalGas(toAddr, fromAddr, amount, path string) []byte {
	var param []interface{}
	gasBody := GasBody{
		Version:    0,
		To:         toAddr,
		From:       fromAddr,
		Value:      amount,
		GasLimit:   0,
		GasFeeCap:  "0",
		GasPremium: "0",
		Method:     0,
		Params:     nil,
	}

	maxFee := MaxFee{MaxFee: "0"}

	param = append(param, gasBody)
	param = append(param, maxFee)
	param = append(param, nil)

	rpc := JsonRpc{
		"2.0",
		"Filecoin.GasEstimateMessageGas",
		param,
		1,
		time.Now().Unix(),
	}

	b, _ := json.Marshal(rpc)

	return b
}

func ServerResponseMarshal(data string, err string) []byte {

	r := Response{Error: err, Data: data}

	b, _ := json.Marshal(r)
	return b
}

func ServerResponseUnMarshal(b []byte) (string, error) {

	r := &Response{}

	_ = json.Unmarshal(b, r)

	return r.Data, errors.New(r.Error)
}
