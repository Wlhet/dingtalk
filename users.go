package dingtalk

import (
	"fmt"
)

//获取企业人数
//@parmas only_active bool   1.false：包含未激活钉钉的人员数量 2.true：只包含激活钉钉的人员数量
func (this *DingTalk) GetUsersNum(only_active bool) int64 {
	token := this.GetToken()
	apiUrl := fmt.Sprintf(GetUsersNumUrl, token)
	PostData := map[string]bool{"only_active": only_active}
	var msg GetUsersNumMsg
	err := PostJsonPtr(apiUrl, PostData, &msg)
	if err == nil {
		return msg.Result.Count
	} else {
		return -1
	}
}

//根据手机号获取userid
//@parmas mobile string
func (this *DingTalk) GetIdByMobile(m string) string {
	token := this.GetToken()
	apiUrl := fmt.Sprintf(GeIdByMobileUrl, token)
	PostData := map[string]string{"mobile": m}
	var msg GeIdByMobileMsg

	err := PostJsonPtr(apiUrl, PostData, &msg)
	if err == nil {
		return msg.Result.UserId
	} else {
		return ""
	}
}

//通过UserId获取用户信息
func (this *DingTalk) GetUserInfo(userid string, language ...string) UserInfoMsg {
	token := this.GetToken()
	apiUrl := fmt.Sprintf(GetUserUrl, token)
	PostData := map[string]string{"userid": userid}
	for _, v := range language {
		if v == "zh_CN" || v == "en_US" {
			PostData["language"] = v
		}
	}
	var msg UserInfoMsg
	PostJsonPtr(apiUrl, PostData, &msg)
	return msg
}

//通过UserId获取用户信息
func (this *DingTalk) UpdateUserInfo(info UserInfo) UpdateUserInfoMsg {
	var msg UpdateUserInfoMsg
	if info.UserId == "" {
		msg.ErrCode = -1000
		msg.ErrMsg = "请传入ID"
		return msg
	}
	token := this.GetToken()
	apiUrl := fmt.Sprintf(UpdateUrl, token)
	PostJsonPtr(apiUrl, info, &msg)
	return msg
}

//新增用户  ID不传的话自动生成
func (this *DingTalk) CreateUser(info UserInfo) CreateUserMsg {
	var msg CreateUserMsg
	token := this.GetToken()
	apiUrl := fmt.Sprintf(CreateUserUrl, token)
	PostJsonPtr(apiUrl, info, &msg)
	return msg
}

//code免登
func (this *DingTalk) UserLoginByCode(code string) UserLoginByCodeMsg {
	token := this.GetToken()
	apiUrl := fmt.Sprintf(GetUserInfoByCodeUrl, token, code)
	var msg UserLoginByCodeMsg
	GetJson(apiUrl, &msg)
	return msg
}
