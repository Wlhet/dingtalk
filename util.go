package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// TimeOut 全局请求超时设置,默认1分钟
var TimeOut time.Duration = 60 * time.Second

// SetTimeOut 设置全局请求超时
func SetTimeOut(d time.Duration) {
	TimeOut = d
}

// httpClient() 带超时的http.Client
func httpClient() *http.Client {
	return &http.Client{Timeout: TimeOut}
}

// GetJson 发送GET请求解析json
func GetJson(uri string, v interface{}) error {

	r, err := httpClient().Get(uri)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

// GetBody 发送GET请求，返回body字节
func GetBody(uri string) ([]byte, error) {
	resp, err := httpClient().Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get err: uri=%v , statusCode=%v", uri, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

//PostJson 发送Json格式的POST请求
func PostJson(uri string, obj interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(obj)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient().Post(uri, "application/json;charset=utf-8", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", uri, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

//PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonPtr(uri string, obj interface{}, result interface{}, contentType ...string) (err error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	//	enc.SetEscapeHTML(false)
	err = enc.Encode(obj)
	if err != nil {
		return
	}
	ct := "application/json;charset=utf-8"
	if len(contentType) > 0 {
		ct = strings.Join(contentType, ";")
	}
	// fmt.Println("post buf:", buf.String()) // Debug
	resp, err := httpClient().Post(uri, ct, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", uri, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(result)
}

//PostJsonPtr 发送Json格式的POST请求并解析结果到result指针
func PostJsonStrPtr(uri string, js string, result interface{}, contentType ...string) (err error) {
	// fmt.Println("post str:", js)
	bufofStr := []byte(js)
	buf := bytes.NewBuffer(bufofStr)
	ct := "application/json;charset=utf-8"
	if len(contentType) > 0 {
		ct = strings.Join(contentType, ";")
	}
	// fmt.Println("post buf:", buf.String()) // Debug
	resp, err := httpClient().Post(uri, ct, buf)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http post error : uri=%v , statusCode=%v", uri, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(result)
}

// Min golang min int
func Min(first int, args ...int) int {
	for _, v := range args {
		if first > v {
			first = v
		}
	}
	return first
}
