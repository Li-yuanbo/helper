package model

import "helper/utils"

type UserInfoModel struct {
	Id         int64  `json:"id"`
	UserName   string `json:"user_name"` //用户名
	Password   string `json:"password"`
	Name       string `json:"name"` //用户昵称
	Phone      string `json:"phone"`
	UserType   int64  `json:"user_type"` //0-普通用户 1-组织管理员 2-admin用户
	Gender     int64  `json:"gender"`
	Age        int64  `json:"age"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type RegisterUserReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	UserType int64  `json:"user_type"`
	Gender   int64  `json:"gender"`
	Age      int64  `json:"age"`
}

type RegisterUserResp struct {
	Res  *utils.Res     `json:"res"`
	User *UserInfoModel `json:"user"`
}

type GetUserInfosReq struct {
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
	CurPage int64 `json:"cur_page"`
}

type GetUserInfosResp struct {
	Res          *utils.Res       `json:"res"`
	CurPage      int64            `json:"cur_page"`
	TotalPage    int64            `json:"total_page"`
	TotalUserNum int64            `json:"total_user_num"`
	Users        []*UserInfoModel `json:"users"`
}
