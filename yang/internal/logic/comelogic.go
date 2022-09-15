package logic

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"service/yang/internal/svc"
	"service/yang/internal/types"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type ComeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewComeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ComeLogic {
	return &ComeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ComeLogic) Come(req *types.ComeReq) (resp *types.ComeResp, err error) {
	// ------------------------------------
	// - 检查是否为空
	// ------------------------------------
	if req.Token == "" {
		return &types.ComeResp{
			Code: 201,
			Msg:  "token is empty",
		}, nil
	}
	// ------------------------------------
	// - request
	// ------------------------------------
	client := &http.Client{}
	httpReq, err := http.NewRequest("GET",
		fmt.Sprintf("https://cat-match.easygame2021.com/sheep/v1/game/game_over?rank_score=1&rank_state=1&rank_time=1&rank_role=1&skin=1&t=%v", req.Token),
		nil)
	if err != nil {
		return &types.ComeResp{
			Code: 202,
			Msg:  "request is error",
		}, nil
	}
	//httpReq.Header.Set("accept", "application/json, text/plain, */*")
	//httpReq.Header.Set("token", "")
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return &types.ComeResp{
			Code: 202,
			Msg:  "request is error",
		}, nil
	}
	bodyText, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return &types.ComeResp{
			Code: 202,
			Msg:  "request is error",
		}, nil
	}
	if strings.Contains(string(bodyText), "没有权限") {
		return &types.ComeResp{
			Code: 203,
			Msg:  "token is invalid",
		}, nil
	}

	return &types.ComeResp{
		Code: 200,
		Msg:  string(bodyText),
	}, nil
}
