package wechatcheckin

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func CheckIn(client *http.Client, headers map[string]string, url string) {
	req, err := http.NewRequest("GET", "https://pass.hust.edu.cn/cas/login", nil)
	if err != nil {
		panic(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if strings.Contains(string(body), "Sign in successfully") {
		fmt.Println("签到成功")
	} else {
		fmt.Println(string(body))
	}
}
