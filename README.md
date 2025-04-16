这是一个 Go 的 华科统一认证登录(http://pass.hust.edu.cn/) 库

## 参考：  
验证码识别:[hustcode](https://github.com/HomeArchbishop/hustcode) (原谅我拙略的模仿，我实现的精度并不高，解决方式是多登录几次)

程序主体:[HustLogin](https://github.com/MarvinTerry/HustLogin)(把二维码识别改为不需要外部OCR了)

最近搜了一下，好像我想干的都有人干过了:  
[GoLoginHust](https://github.com/black-binary/GoLoginHust)  
[libhustpass](https://github.com/naivekun/libhustpass)  
[Sign your horse 签个马](https://github.com/naivekun/sign-your-horse)

不过写都写了就放出来吧，有问题欢迎提ISSUE  
（本人是 go 新手，代码写得💩请见谅，欢迎拷打）

## 目录结构
解释一下目录结构（其实没什么结构，随便写的）：
- codeGifs： 测试用的验证码 gif
- datas: 拿来存账号密码的，如果你想的话
- getWork: 一个查询劳动课的 demo
- header: 存放一些请求头
- login: 登录主要文件
- newGetCode: 识别验证码
- outputImg: 测试用的图片处理
- util: 工具函数
- WechatCheckIn: 企业微信二维码签到（超星学习通），目前还没测试过，不知道能不能用
- withLogin: 对 login 的封装  

- work.go: 获取劳动课并推送通知
---
## 最佳实践：
使用`withLogin`中的相关方法，请移步[./withLogin/README.md](./withLogin/README.md)  

具体实践，参考 [./work.go](./work.go) 
