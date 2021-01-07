package safecustody_filecoin

import "github.com/filecoin-project/go-address"

type Hex struct {
	Hex string `json:"hex"`
	T   int64  `json:"t"`
}

type Nonce struct {
	FromAddr address.Address `json:"from_addr"`
	T        int64           `json:"t"`
}

type GasBody struct {
	Version    float64 `json:"Version"`
	To         string  `json:"To"`
	From       string  `json:"From"`
	Nonce      int64   `json:"Nonce"`
	Value      string  `json:"Value"`
	GasLimit   int64   `json:"GasLimit"`
	GasFeeCap  string  `json:"GasFeeCap"`
	GasPremium string  `json:"GasPremium"`
	Method     int     `json:"Method"`
	Params     []byte  `json:"Params"`
}

type MaxFee struct {
	MaxFee string `json:"MaxFee"`
}

type JsonRpc struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
	T       int64       `json:"t"`
}

type Response struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}

type RespGas struct {
	JsonRpc string  `json:"jsonrpc"`
	Result  GasBody `json:"result"`
}
