package utils

var ERR_CODE = map[int]string{
	//通用错误
	10001: "GET_REQUEST_BODY_ERROR", //获取request的body错误
	10002: "DATA_UNMARSHAL_ERROR",   //json反序列化错误
	10003: "SQL_ERROR",              //数据库获取错误
	10004: "PARAM_ERROR",            //参数错误
	10005: "REQUEST_IS_NIL",         //request body为空
	10006: "SESSION_SAVE_ERR",       //保存session错误

	//user_info
	20001: "USER_PASSWORD_ERROR", //用户名错误
	20002: "USER_USERNAME_ERROR", //用户名错误或未注册
	20003: "USER_ALREADY_LOGIN",  //用户已登录
	20004: "USER_NOT_LOGIN_ERR",  //用户未登录

	//org
	30001: "ORG_USERNAME_ERROR",    //组织账号错误
	30002: "ORG_PASSWORD_ERROR",    //组织密码错误
	30003: "ORG_ALREADY_LOGIN_ERR", //机构已经登录
	30004: "ORG_NOT_LOGIN_ERR",     //机构已经登录
}

type Res = struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
