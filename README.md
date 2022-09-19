# 羊了个羊在线刷榜通关辅助小工具

# API地址 
```
/chabai/v1
req {
    token: string
}
```
# web地址
`/assets/index.htm`



# 成品启动方式

win
`./yang.exe -f etc/chabai-api.yaml`

linux
`./yang -f etc/chabai-api.yaml`


# 源码启动方式

`go run yang.go -f etc/chabai-api.yaml`

# Token获取方式

PC可以下载软件或者易语言源码
Android IOS需使用抓包软件
Android 7.0还是几点0以上好像没办法安装证书 就没办法解密 https

# 免token uid模式
```
// 1.获取openid - get请求
https://cat-match.easygame2021.com/sheep/v1/game/user_info?uid=66666666&t=token
// 2.拿到openid - post json
https://cat-match.easygame2021.com/sheep/v1/user/login_oppo
{
	"uid": "上面获取的openid",
	"nick_name": "nickName",
	"avatar": "avatar",
	"sex": 1
}
```
