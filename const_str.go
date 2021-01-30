package dingtalk

const (
	RegCallBackUrl              = "https://oapi.dingtalk.com/call_back/register_call_back?access_token=%s" //注册
	UpdateCallBackUrl           = "https://oapi.dingtalk.com/call_back/update_call_back?access_token=%s"   //更新
	GetCallBackUrl              = "https://oapi.dingtalk.com/call_back/get_call_back?access_token=%s"      //查询
	DeleteCallBackUrl           = "https://oapi.dingtalk.com/call_back/delete_call_back?access_token=%s"
	GetTokenUrl                 = "https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s"
	SendMsgToUserUrl            = "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s"
	processinstanceUrl          = "https://oapi.dingtalk.com/topapi/processinstance/get?access_token=%s"        //获取审批实例地址
	GetProNameUrl               = "https://oapi.dingtalk.com/topapi/process/get_by_name?access_token=%s"        //获取审批Id
	ProcessinstanceTerminateUrl = "https://oapi.dingtalk.com/topapi/process/instance/terminate?access_token=%s" //终止实例接口地址
	GetUsersNumUrl              = "https://oapi.dingtalk.com/topapi/user/count?access_token=%s"
	GeIdByMobileUrl             = "https://oapi.dingtalk.com/topapi/v2/user/getbymobile?access_token=%s"
	GetUserUrl                  = "https://oapi.dingtalk.com/topapi/v2/user/get?access_token=%s"
	UpdateUrl                   = "https://oapi.dingtalk.com/topapi/v2/user/update?access_token=%s"
	CreateUserUrl               = "https://oapi.dingtalk.com/topapi/v2/user/create?access_token=%s"
	GetUserInfoByCodeUrl        = "https://oapi.dingtalk.com/user/getuserinfo?access_token=%s&code=%s"
)
