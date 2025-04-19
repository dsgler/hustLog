package withlogin

import (
	"crypto/tls"
	"errors"
	"hustLog/header"
	"hustLog/login"
	"hustLog/util"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

var ErrMaxRetryExceeded = errors.New("登录失败，超过最大重试次数，退出程序")

func New(rawU string, rawP string, headers map[string]string, maxRetry int, proxyURL *url.URL) (*WithLogin, error) {
	var client *http.Client
	retryCnt := 0
	var wl = &WithLogin{Client: client, DefaultHeaders: headers, User: rawU}
	for retryCnt < maxRetry {
		var err error
		client, err = login.HustLogin(rawU, rawP, headers, proxyURL)
		if err != nil {
			log.Println(err)
			retryCnt++
			continue
		}
		wl.Client = client

		if name, err := wl.CheckLogin(); err != nil || name == "" {
			log.Println(err)
			retryCnt++
			continue
		}

		break
	}

	if retryCnt == maxRetry {
		return nil, ErrMaxRetryExceeded
	}

	return wl, nil
}

type WithLogin struct {
	Client         *http.Client
	DefaultHeaders map[string]string
	User           string
}

func (w *WithLogin) Get(url string, headers map[string]string) (*http.Response, []byte, error) {
	if headers == nil {
		headers = w.DefaultHeaders
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := w.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	return resp, body, nil

}

func (w *WithLogin) Post(url string, body io.Reader, headers map[string]string) (*http.Response, []byte, error) {
	if headers == nil {
		headers = w.DefaultHeaders
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := w.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	rbody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	return resp, rbody, nil

}

func (w *WithLogin) CheckLogin() (string, error) {
	body, err := util.Myget(w.Client, "https://one.hust.edu.cn/dcp/forward.action?path=/portal/portal&p=home", header.Headers)
	if err != nil {
		return "", err
	}

	rest := regexp.MustCompile(`usernameandidnumber="([^"]+)"`).FindStringSubmatch(string(body))
	if len(rest) < 2 {
		log.Println("登陆失败")
		log.Println(string(body))
		return "", nil
	} else {
		name := rest[1]
		log.Println("登录成功,欢迎：" + name)
		return name, nil
	}
}

func (w *WithLogin) SetProxy(proxyurl string, isInsecure bool) error {
	u, err := url.Parse(proxyurl)
	if err != nil {
		return err
	}
	if isInsecure {
		w.Client.Transport = &http.Transport{
			Proxy:           http.ProxyURL(u),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		w.Client.Transport = &http.Transport{
			Proxy:           http.ProxyURL(u),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}
	}
	return nil
}
