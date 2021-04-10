package model

type UserInfoModel struct {
	Id			int64	`json:"id"`
	UserName	string  `json:"user_name"`	//用户名
	Password	string  `json:"password"`
	Name		string  `json:"name"`	//用户昵称
	Phone		string  `json:"phone"`
	UserType	int64   `json:"user_type"`	//0-普通用户 1-admin用户
	Gender		int64   `json:"gender"`
	Age			int64   `json:"age"`
	CreateTime	int64   `json:"create_time"`
	UpdateTime	int64   `json:"update_time"`
}

type RegisterUserReq struct {
	UserName	string  `json:"user_name"`
	Password	string  `json:"password"`
	Name		string  `json:"name"`
	Phone		string  `json:"phone"`
	UserType	int64   `json:"user_type"`
	Gender		int64   `json:"gender"`
	Age			int64   `json:"age"`
}

type RegisterUserResp struct {
	User *UserInfoModel `json:"user"`
}