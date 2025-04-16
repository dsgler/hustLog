package login

import "net/http"

func NoRedirect(req *http.Request, via []*http.Request) error {
	// 不跟随重定向
	return http.ErrUseLastResponse
}

type publicKeyStruct struct {
	Key string `json:"publicKey"`
}
