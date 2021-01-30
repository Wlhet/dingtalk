package dingtalk

import (
	"fmt"
)

//https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-details-about-a-specified-approval-instance
type ProInstance struct {
	Msg
	RequestId       string `json:"request_id"`
	ProcessInstance struct {
		Title                      string   `json:"title"`
		CreateTime                 int64    `json:"create_time"`
		FinishTime                 int64    `json:"finish_time"`
		OriginatorUserId           string   `json:"originator_userid"`             //发起人ID
		OriginatorDeptId           string   `json:"originator_dept_id"`            //发起人部门   -1跟部门
		Status                     string   `json:"status"`                        //状态码 NEW：新创建 RUNNING：审批中 TERMINATED：被终止 COMPLETED：完成 CANCELED：取消
		ApproverUserids            []string `json:"approver_userids"`              //审批人ID素组
		CcUserids                  []string `json:"cc_userids"`                    //抄送人ID数组
		Result                     string   `json:"result"`                        //结果
		BusinessId                 string   `json:"business_id"`                   //审批实例业务编号。
		AttachedProcessInstanceIds []string `json:"attached_process_instance_ids"` //附属实例
		BizAction                  string   `json:"biz_action"`                    //MODIFY：表示该审批实例是基于原来的实例修改而来 REVOKE：表示该审批实例是由原来的实例撤销后重新发起的 NONE表示正常发起
		OriginatorDeptName         string   `json:"originator_dept_name"`          //发起部门
		FormComponentValues        []struct {
			Name          string `json:"name"`
			Value         string `json:"value"`
			Id            string `json:"id"`
			ComponentType string `json:"component_type"`
		} `json:"form_component_values"`
	} `json:"process_instance"`
}

type ProcessCodeMsg struct {
	Msg
	ProcessCode string `json:"process_code"`
}

func (this *DingTalk) GetProInstance(pid string) ProInstance {
	token := this.GetToken()
	apiUrl := fmt.Sprintf(processinstanceUrl, token)
	PostData := map[string]string{"process_instance_id": pid}
	var msg ProInstance
	err := PostJsonPtr(apiUrl, PostData, &msg)
	if err == nil {
		return msg
	} else {
		return msg
	}
}
func (this *DingTalk) GetProInstance1(pid string) string {
	token := this.GetToken()
	apiUrl := fmt.Sprintf(processinstanceUrl, token)
	PostData := map[string]string{"process_instance_id": pid}
	rs, err := PostJson(apiUrl, PostData)
	if err == nil {
		return string(rs)
	} else {
		return string(rs)
	}
}

func (this *DingTalk) GetBpmProcessCodeByName(name string) string {
	token := this.GetToken()
	apiUrl := fmt.Sprintf(ProcessinstanceTerminateUrl, token)
	PostData := map[string]string{"name": name}
	var msg ProcessCodeMsg
	err := PostJsonPtr(apiUrl, PostData, &msg)
	if err == nil {
		return msg.ProcessCode
	} else {
		return ""
	}
}

func (this *DingTalk) StopProInstance(pid string, remark string) StopProInstanceMsg {

	token := this.GetToken()
	apiUrl := fmt.Sprintf(ProcessinstanceTerminateUrl, token)

	PostData := map[string]interface{}{
		"request": map[string]interface{}{
			"is_system":           true,
			"process_instance_id": pid,
			"remark":              remark,
		},
	}
	var msg StopProInstanceMsg

	err := PostJsonPtr(apiUrl, PostData, &msg)
	// fmt.Printf("%-v", PostData)

	if err == nil {

		return msg
	} else {
		fmt.Println(4, err.Error())
		return msg
	}

}
