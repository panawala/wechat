package custom

import (
	"encoding/json"
	"fmt"

	"github.com/panawala/wechat/context"
	"github.com/panawala/wechat/util"
	"github.com/panawala/wechat/message"
)

const (
	sendCustomURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

//Custom 素材管理
type Custom struct {
	*context.Context
}

//NewCustom init
func NewCustom(context *context.Context) *Custom {
	custom := new(Custom)
	custom.Context = context
	return custom
}

//resAddCustom 永久性素材上传返回的结果
type resAddCustom struct {
	util.CommonError
}

//SendMessage 发送客服消息
func (custom *Custom) SendMessage(toUser string, message *message.CustomMessage) (err error) {
	message.SetToUser(toUser)
	var accessToken string
	accessToken, err = custom.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s&type=video", sendCustomURL, accessToken)

	var response []byte
	response, err = util.PostJSON(uri, message)
	if err != nil {
		return
	}

	var resAddCustom resAddCustom
	err = json.Unmarshal(response, &resAddCustom)
	if err != nil {
		return
	}
	if resAddCustom.ErrCode != 0 {
		err = fmt.Errorf("AddCustom error : errcode=%v , errmsg=%v", resAddCustom.ErrCode, resAddCustom.ErrMsg)
		return
	}
	return
}
