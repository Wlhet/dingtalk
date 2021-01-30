package dingtalk

import (
	"errors"
	"fmt"
)

//{"msgtype":"text","text":{"content":"请提交日报。"}
func (this *DingTalk) SendMsgToUser(userid, msg string) error {
	token := this.GetToken()
	apiUrl := fmt.Sprintf(SendMsgToUserUrl, token)
	PostData := map[string]interface{}{
		"msg": map[string]interface{}{
			"msgtype": "text",
			"text": map[string]string{
				"content": msg,
			},
		},
		"agent_id":    "859987395",
		"userid_list": userid,
	}
	var backmsg Msg
	err := PostJsonPtr(apiUrl, PostData, &backmsg)
	if err == nil {
		if backmsg.ErrCode == 0 {
			return nil
		} else {
			return errors.New(backmsg.ErrMsg)
		}
	} else {
		return errors.New(backmsg.ErrMsg)
	}
}
