package dingtalk

import "fmt"

type CallBack struct {
	CallBackTag []string `json:"call_back_tag"`
	Token       string   `json:"token"`
	AesKey      string   `json:"aes_key"`
	Url         string   `json:"url"`
}

type CallBackResponse struct {
	CallBack
	Msg
}

//注册回调
func (this *DingTalk) RegCallBack(cb CallBack) Msg {
	access_token := this.GetToken()
	apiUrl := fmt.Sprintf(RegCallBackUrl, access_token)
	var msg Msg
	PostJsonPtr(apiUrl, cb, &msg)
	return msg
}

//更新回调
func (this *DingTalk) UpdateCallBack(cb CallBack) Msg {
	access_token := this.GetToken()
	apiUrl := fmt.Sprintf(UpdateCallBackUrl, access_token)
	var msg Msg
	PostJsonPtr(apiUrl, cb, &msg)
	return msg
}

//获取回调
func (this *DingTalk) GetCallBack() CallBackResponse {
	var cb CallBackResponse
	access_token := this.GetToken()
	apiUrl := fmt.Sprintf(GetCallBackUrl, access_token)
	GetJson(apiUrl, &cb)
	return cb
}

//删除回调
func (this *DingTalk) DeleteCallBack() Msg {
	var msg Msg
	access_token := this.GetToken()
	apiUrl := fmt.Sprintf(DeleteCallBackUrl, access_token)
	GetJson(apiUrl, &msg)
	return msg
}
