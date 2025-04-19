package withlogin

import (
	"crypto/tls"
	"encoding/json"
	"hustLog/header"
	"hustLog/login"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"
)

type CookieStore struct {
	C1   []*http.Cookie
	C2   []*http.Cookie
	User string
}

var u1, _ = url.Parse("https://pass.hust.edu.cn")
var u2, _ = url.Parse("https://pass.hust.edu.cn/cas/")

func (w *WithLogin) StoreCookie(path string) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()

	// 获取所有 Cookie

	data, err := json.Marshal(CookieStore{C1: w.Client.Jar.Cookies(u1), C2: w.Client.Jar.Cookies(u2), User: w.User})
	if err != nil {
		return err
	}

	_, err = fp.Write(data)
	return err
}

func (w *WithLogin) LoadCookie(path string) error {
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	var cs *CookieStore
	err = json.NewDecoder(fp).Decode(&cs)
	if err != nil {
		return err
	}

	w.Client.Jar.SetCookies(u1, cs.C1)
	w.Client.Jar.SetCookies(u2, cs.C2)
	return nil
}

// 持久化，避免每次都要登录
func LoadOrNew(rawU string, rawP string, headers map[string]string, maxRetry int, proxyURL *url.URL, cookiePath string) (*WithLogin, error) {
	if rawU == "" || rawP == "" {
		return nil, login.ErrNoUsernameOrPasswd
	}

	if headers == nil {
		headers = header.Headers
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	var tp http.RoundTripper
	if proxyURL == nil {
		tp = &http.Transport{
			Proxy:           http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		tp = nil
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
		// 设置代理
		Transport: tp,
		Jar:       jar,
	}

	wl := &WithLogin{Client: client, DefaultHeaders: headers, User: rawU}
	wl.LoadCookie(cookiePath)

	name, err := wl.CheckLogin()
	if err != nil || name == "" {
		log.Println("载入失败", err)
		wl, err := New(rawU, rawP, headers, maxRetry, proxyURL)
		if err != nil {
			return nil, err
		}
		wl.StoreCookie(cookiePath)
		return wl, nil
	}

	log.Println("载入成功")
	return wl, nil
}
