package dingtalk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDingtalkByJson(CorpId, AppKey, AppSecret string) DingTalk {
	data, err := ioutil.ReadFile("token.json")
	dt := DingTalk{CorpId: CorpId, AppKey: AppKey, AppSecret: AppSecret}
	var tk Token
	if err == nil {
		json.Unmarshal(data, &tk)
		dt.AccessToken = tk.AccessToken
		dt.ExpiresIn = tk.ExpiresIn
	}
	return dt
}

func NewDingtalk(AppKey, AppSecret string) DingTalk {
	return DingTalk{AppKey: AppKey, AppSecret: AppSecret}
}

func (this *DingTalk) GetToken() string {
	nowTime := time.Now().Unix()
	if nowTime > this.ExpiresIn {
		err := this.getToken()
		if err != nil {
			fmt.Print("Token  重新获取失败\n")
			return ""
		} else {
			this.WriteToFile()
			return this.AccessToken
		}
	} else {
		return this.AccessToken
	}
}

//发送网络请求
func (this *DingTalk) getToken() error {
	client := &http.Client{}
	var msg AccessTokenMsg
	FullUrl := fmt.Sprintf(GetTokenUrl, this.AppKey, this.AppSecret)
	request, _ := http.NewRequest("GET", FullUrl, nil)
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(body, &msg)
	nowTime := time.Now().Unix()
	this.ExpiresIn = msg.ExpiresIn + nowTime
	this.AccessToken = msg.AccessToken
	return nil
}

//token存入文件缓存
func (this *DingTalk) WriteToFile() {
	tk := Token{this.AccessToken, this.ExpiresIn}
	data, _ := json.Marshal(tk)
	ioutil.WriteFile("token.json", data, 0666)
}
