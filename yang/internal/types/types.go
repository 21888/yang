// Code generated by goctl. DO NOT EDIT.
package types

type ComeReq struct {
	Token string `json:"token"`
}

type ComeResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}