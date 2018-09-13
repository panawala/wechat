package qrcode

import (
	"encoding/json"
	"fmt"

	"github.com/panawala/wechat/context"
	"github.com/panawala/wechat/util"
)

const (
	addQrcodeURL = "https://api.weixin.qq.com/cgi-bin/qrcode/create"
)

// Qrcode ActionName 类型
type ActionName string

const (
	//临时的整型参数值
	QR_SCENE ActionName = "QR_SCENE"
	//临时的字符串参数值
	QR_STR_SCENE = "QR_STR_SCENE"
	//永久的整型参数值
	QR_LIMIT_SCENE = "QR_LIMIT_SCENE"
	//永久的字符串参数值
	QR_LIMIT_STR_SCENE = "QR_LIMIT_STR_SCENE"
)

//Qrcode 二维码管理
type Qrcode struct {
	*context.Context
}

//NewQrcode init
func NewQrcode(context *context.Context) *Qrcode {
	qrcode := new(Qrcode)
	qrcode.Context = context
	return qrcode
}

//reqQrcode 请求结果

type actionInfo struct {
	SceneId  int    `json:"scene_id,omitempty"`
	SceneStr string `json:"scene_str,omitempty"`
}

type reqQrcode struct {
	ExpireSeconds int        `json:"expire_seconds,omitempty"`
	ActionName    string     `json:"action_name"`
	ActionInfo    actionInfo `json:"action_info"`
}

//resQrcode 返回结果
type resQrcode struct {
	util.CommonError

	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	Url           string `json:"url"`
}

//AddQrcode 上传二维码
func (qrcode *Qrcode) AddQrcode(actionName string, expireSeconds, sceneId int, sceneStr string) (ticket string, qrCodeUrl string, err error) {
	ticket = ""
	qrCodeUrl = ""
	var accessToken string
	accessToken, err = qrcode.GetAccessToken()
	if err != nil {
		return
	}

	req := &reqQrcode{
		ExpireSeconds: expireSeconds,
		ActionName:    actionName,
		ActionInfo:    actionInfo{SceneId: sceneId, SceneStr: sceneStr},
	}

	uri := fmt.Sprintf("%s?access_token=%s", addQrcodeURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}
	var resQrcode resQrcode
	err = json.Unmarshal(response, &resQrcode)
	if err != nil {
		return
	}
	if resQrcode.ErrCode != 0 {
		err = fmt.Errorf("AddQrcode error : errcode=%v , errmsg=%v", resQrcode.ErrCode, resQrcode.ErrMsg)
		return
	}
	ticket = resQrcode.Ticket
	qrCodeUrl = resQrcode.Url
	return
}
