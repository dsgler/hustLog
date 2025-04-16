package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"hustLog/datas"
	getwork "hustLog/getWork"
	"hustLog/header"
	withlogin "hustLog/withLogin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var path, _ = os.Getwd()

var cookiePath = filepath.Join(path, "cookie.json")
var kPath = filepath.Join(path, "oldKXIDs.json")

type sendM struct {
	Content string   `json:"content"`
	SptList []string `json:"sptList"`
}

func main() {
	wl, err := withlogin.LoadOrNew(datas.Username, datas.Password, header.Headers, 3, nil, cookiePath)
	if err != nil {
		panic(err)
	}
	body := getwork.MustQueryWork1(wl)

	message, isAnyNew := getwork.FilterAvailible(wl, body, kPath)
	log.Println(message)

	if isAnyNew {
		send := bytes.NewBuffer(nil)
		// 通过wxPusher推送
		json.NewEncoder(send).Encode(&sendM{Content: message, SptList: []string{"your SPT"}})
		// log.Println(send.String())

		client := &http.Client{
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		}

		resp, err := client.Post("https://wxpusher.zjiecode.com/api/send/message/simple-push", "application/json", send)
		if err != nil {
			log.Fatalln(err)
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		log.Println(string(body))
	}

}
