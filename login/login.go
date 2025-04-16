package login

import (
	"bytes"
	"crypto/tls"
	"strconv"

	"errors"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"

	newgetcode "hustLog/newGetCode"

	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"hustLog/header"
	"hustLog/util"
)

var ErrNoUsernameOrPasswd = errors.New("请输入 用户名 或 密码")

func HustLogin(rawU, rawP string, headers map[string]string, proxyURL *url.URL) (*http.Client, error) {
	if rawU == "" || rawP == "" {
		return nil, ErrNoUsernameOrPasswd
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
		Timeout: 5 * time.Second,
		// 设置代理
		Transport:     tp,
		CheckRedirect: NoRedirect,
		Jar:           jar,
	}

	// 恢复重定向
	defer func() { client.CheckRedirect = nil }()

	// 开始获取密码
	body, err := util.Myget(client, "https://pass.hust.edu.cn/cas/login", headers)
	if err != nil {
		return nil, err
	}

	img, err := util.Myget(client, "https://pass.hust.edu.cn/cas/code", headers)
	if err != nil {
		return nil, err
	}
	code, err := newgetcode.MergeAndGet(img)
	if err != nil {
		return nil, err
	}
	form := regexp.MustCompile(`(?s)<form id="loginForm" (.*)</form>`).FindString(string(body))
	nonce := regexp.MustCompile(`<input type="hidden" id="lt" name="lt" value="(.*)" />`).FindStringSubmatch(form)[1]
	execution := regexp.MustCompile(`<input type="hidden" name="execution" value="(.*)" />`).FindStringSubmatch(form)[1]

	// 开始加密
	// 获取rsa，这是必须的
	req, _ := http.NewRequest("POST", "https://pass.hust.edu.cn/cas/rsa", nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, _ := client.Do(req)
	body, err = io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	publicKey := publicKeyStruct{}
	err = json.Unmarshal(body, &publicKey)
	if err != nil {
		panic(err)
	}
	base64PublicKey := publicKey.Key

	encryptedU := EncodeRSA(rawU, base64PublicKey)
	encryptedP := EncodeRSA(rawP, base64PublicKey)
	postParams := map[string]string{
		"rsa":       "",
		"ul":        encryptedU,
		"pl":        encryptedP,
		"code":      code,
		"phoneCode": "",
		"lt":        nonce,
		"execution": execution,
		"_eventId":  "submit",
	}

	postData := encodeMap(postParams)

	req, err = http.NewRequest("POST", "https://pass.hust.edu.cn/cas/login", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusFound {
		return nil, errors.New("---HustPass Failed---" + "状态为:" + strconv.Itoa(resp.StatusCode) + resp.Status)
	}

	log.Println("---HustPass Succeed---")
	return client, nil
}

func encodeMap(ma map[string]string) (str string) {
	strs := make([]string, 0, 10)
	// 不清楚顺序是否重要
	keys := []string{"ul", "pl", "code", "lt", "execution", "_eventId"}
	for i := 0; i < len(keys); i++ {
		k := keys[i]
		v := ma[k]
		s := k + "=" + url.QueryEscape(v)
		// s = url.QueryEscape(s)
		strs = append(strs, s)
	}
	str = strings.Join(strs, "&")
	return
}

func EncodeRSA(str string, base64PublicKey string) string {
	derPublicKey, err := base64.StdEncoding.DecodeString(base64PublicKey)
	if err != nil {
		panic(err)
	}
	rpublicKey, err := x509.ParsePKIXPublicKey(derPublicKey)
	publicKey := rpublicKey.(*rsa.PublicKey)
	if err != nil {
		panic(err)
	}
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(str))
	if err != nil {
		panic(err)
	}
	ret := base64.StdEncoding.EncodeToString(ciphertext)
	return ret
}

func CheckLogin(c *http.Client) bool {
	body, err := util.Myget(c, "https://one.hust.edu.cn/dcp/forward.action?path=/portal/portal&p=home", header.Headers)
	if err != nil {
		log.Print("登陆失败")
		log.Println(err)
		return false
	}

	rest := regexp.MustCompile(`usernameandidnumber="([^"]+)"`).FindStringSubmatch(string(body))
	if len(rest) < 2 {
		log.Println("登陆失败\n" + string(body))
		return false
	} else {
		log.Println("登录成功,欢迎：" + rest[1])
		return true
	}
}
