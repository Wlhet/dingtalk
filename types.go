package dingtalk

type DingTalk struct {
	CorpId      string
	AppKey      string //账号
	AppSecret   string //密钥
	AccessToken string //token
	ExpiresIn   int64  //过期时间
}
type Msg struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

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
type GeIdByMobileMsg struct {
	Msg
	RequestId string `json:"request_id"`
	Result    struct {
		UserId string `json:"userid"`
	} `json:"result"`
}

type UserInfoMsg struct {
	Msg
	UserInfo `json:"result"`
}

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

type UpdateUserInfoMsg struct {
	Msg
}
type CreateUserMsg struct {
	Msg
	Result struct {
		UserId string `json:"userid"`
	} `json:"result"`
}

type StopProInstanceMsg struct {
	Msg
	Result  bool `json:"result"`  //终止成功
	Success bool `json:"success"` //调用成功
}

type UserLoginByCodeMsg struct {
	Msg
	UserId   string `json:"userId"`
	Name     string `json:"name"`
	DeviceId string `json:"deviceId"`
	IsSys    bool   `json:"is_sys"`
	SysLevel int64  `json:"sys_level"`
}
