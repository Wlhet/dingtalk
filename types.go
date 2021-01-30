package dingtalk

import "crypto/cipher"

type DingTalk struct {
	CorpId    string
	AppKey    string //账号
	AppSecret string //密钥
	Token
}
type Msg struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
type Crypto struct {
	Token    string
	AesKey   string
	SuiteKey string
	block    cipher.Block
	bkey     []byte
}

//回调信息
type CallBack struct {
	CallBackTag []string `json:"call_back_tag"`
	Token       string   `json:"token"`
	AesKey      string   `json:"aes_key"`
	Url         string   `json:"url"`
}

//回调信息返回消息
type CallBackResponse struct {
	CallBack
	Msg
}

//审批时间结构体
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

//获取token返回信息
type AccessTokenMsg struct {
	Msg
	Token
}

type GetUsersNumMsg struct {
	Msg
	RequestId string `json:"request_id"`
	Result    struct {
		Count int64 `json:"count"`
	} `json:"result"`
}

//手机号获取用户Id返回信息
type GeIdByMobileMsg struct {
	Msg
	RequestId string `json:"request_id"`
	Result    struct {
		UserId string `json:"userid"`
	} `json:"result"`
}

//获取用户基本信息返回消息
type UserInfoMsg struct {
	Msg
	UserInfo `json:"result"`
}

//用户基本信息
type UserInfo struct {
	UserId       string     `json:"userid"`
	Unionid      string     `json:"unionid,omitempty"`
	Name         string     `json:"name,omitempty"`
	Avatar       string     `json:"avatar,omitempty"` //头像
	Mobile       string     `json:"mobile,omitempty"`
	JobNumber    string     `json:"job_number,omitempty"` //工号
	Title        string     `json:"title,omitempty"`      //职位
	Email        string     `json:"email,omitempty"`
	Remark       string     `json:"remark,omitempty"`
	DeptIdList   []int64    `json:"dept_id_list,omitempty"`
	Extension    string     `json:"extension,omitempty"`   //爱好 {"爱好":"旅游","年龄":"24"}
	HiredDate    int64      `json:"hired_date,omitempty"`  //入职时间戳
	Active       bool       `json:"active,omitempty"`      //是否激活钉钉
	RealAuthed   bool       `json:"real_authed,omitempty"` //是否实名
	Senior       bool       `json:"senior,omitempty"`      //是否高管
	Admin        bool       `json:"admin,omitempty"`       //是否管理员
	Boss         bool       `json:"boss,omitempty"`        //是否老板
	LeaderInDept []struct { //员工在对应的部门中是否领导
		DeptId int  `json:"dept_id,omitempty"`
		Leader bool `json:"leader,omitempty"`
	} `json:"leader_in_dept,omitempty"`
	RoleList []struct {
		Id        int    `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		GroupName string `json:"group_name,omitempty"`
	} `json:"role_list,omitempty"`
}

//更新用户返回信息
type UpdateUserInfoMsg struct {
	Msg
}

//新建用户返回信息
type CreateUserMsg struct {
	Msg
	Result struct {
		UserId string `json:"userid"`
	} `json:"result"`
}

//终止审批实例返回信息
type StopProInstanceMsg struct {
	Msg
	Result  bool `json:"result"`  //终止成功
	Success bool `json:"success"` //调用成功
}

//用户登录Code
type UserLoginByCodeMsg struct {
	Msg
	UserId   string `json:"userId"`
	Name     string `json:"name"`
	DeviceId string `json:"deviceId"`
	IsSys    bool   `json:"is_sys"`
	SysLevel int64  `json:"sys_level"`
}
