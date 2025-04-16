这是对 login 的封装

主要函数有：
```go
func New(rawU string, rawP string, headers map[string]string, maxRetry int, proxyURL *url.URL) (*WithLogin, error)
```
参数都很显然，`rawU`指账号，`rawP`指密码  
`header`可以用[`header.Headers`](../header/header.go)，或者是不同网站抓包获得  
`maxRetry`重试次数，主要因为二维码识别概率出错，建议为3  
`proxyURL`代理，没有需要直接传`nil`

```go
func (w *WithLogin) Get(url string, headers map[string]string) (*http.Response, []byte, error)
```
发送 GET 请求，`headers`为空将使用默认header，第二个返回值为body

```go
func (w *WithLogin) Post(url string, body io.Reader, headers map[string]string) (*http.Response, []byte, error)
```
发送 POST 请求，`headers`为空将使用默认header，第二个返回值为body

```go
func (w *WithLogin) CheckLogin() (bool, error)
```
检查是否处于登录状态

---
除此之外，在[storage.go](./storage.go)文件内还有保存cookie的方法（数据持久化，避免每次都要登录）
```go
func LoadOrNew(rawU string, rawP string, headers map[string]string, maxRetry int, proxyURL *url.URL, cookiePath string) (*WithLogin, error)
```
从`cookiePath`中读取 cookie 并检查登录状态，如果失败将调用 New 新登录，并将 cookie 保存在`cookiePath`中
