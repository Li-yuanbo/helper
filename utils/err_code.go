package utils

var(
	ERR_CODE = map[int]string{
		//通用错误
		10001: "GET_REQUEST_BODY_ERROR",	//获取request的body错误
		10002: "DATA_UNMARSHAL_ERROR",		//json反序列化错误
		10003: "SQL_ERROR",					//数据库获取错误
		10004: "PARAM_ERROR",				//参数错误
		10005: "REQUEST_IS_NIL",			//request body为空
	}
)