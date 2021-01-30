package dingtalk

type BpmsEvent struct {
	EventType         string `json:"EventType"`         //事件类型
	ProcessInstanceId string `json:"processInstanceId"` //审批实例id
	CorpId            string `json:"corpId"`            //审批实例对应的企业
	CreateTime        int64  `json:"createTime"`        //实例创建时间。
	Title             string `json:"title"`             //实例标题
	Type              string `json:"type"`              //类型，type为start表示审批实例开始 /finish：审批正常结束（同意或拒绝 terminate：审批终止（发起人撤销审批单）
	StaffId           string `json:"staffId"`           //发起审批实例的员工
	Url               string `json:"url"`               //审批实例url，可在钉钉内跳转到审批页面
	BizCategoryId     string `json:"bizCategoryId"`     //审批实例对应表单类别
	FinishTime        int64  `json:"finishTime"`        //审批结束时间
	Result            string `json:"result"`            //redirect
	Remark            string `json:"remark"`            //remark表示操作时写的评论内容
	TaskId            int64  `json:"taskId"`            //任务ID
	Content           string `json:"content"`           //内容
	ProcessCode       string `json:"processCode"`       //审批模板的唯一码
	BusinessId        string `json:"businessId"`        //咱不知道
}
