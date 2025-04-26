package ancheck

import (
	"bytes"
	"encoding/json"
	"hustLog/header"
	withlogin "hustLog/withLogin"
	"maps"
)

type MyClient struct {
	wl      *withlogin.WithLogin
	IsLoged bool
	User    string
}

var userStore map[string]*MyClient

func InitStore() {
	userStore = map[string]*MyClient{}
}

func SetStore(user string, mc *MyClient) {
	userStore[user] = mc
}

func GetStore(user string) *MyClient {
	return userStore[user]
}

func (mc *MyClient) Login(rawU string, rawP string, maxRetry int) (err error) {
	wl, err := withlogin.New(rawU, rawP, header.WechatHeader, 3, nil)
	if err != nil {
		return
	}

	mc.wl = wl
	mc.User = wl.User
	return
}

func (mc *MyClient) CheckIsLoged() (name string, err error) {
	name, err = mc.wl.CheckLogin()
	return
}

type NetRet struct {
	StatusCode int
	Body       []byte
}

func (mc *MyClient) Get(url string, headerString string) (ret *NetRet, err error) {
	var headers header.HeaderType
	json.Unmarshal([]byte(headerString), &headers)
	maps.Copy(headers, header.WechatHeader)

	res, body, err := mc.wl.Get(url, headers)
	if err != nil {
		return
	}

	ret = &NetRet{StatusCode: res.StatusCode, Body: body}
	return
}

func (mc *MyClient) Post(url string, body []byte, headerString string) (ret *NetRet, err error) {
	var headers header.HeaderType
	json.Unmarshal([]byte(headerString), &headers)
	maps.Copy(headers, header.WechatHeader)

	res, body, err := mc.wl.Post(url, bytes.NewReader(body), headers)
	if err != nil {
		return
	}

	ret = &NetRet{StatusCode: res.StatusCode, Body: body}
	return
}

func Bytes2String(b []byte) string {
	return string(b)
}
