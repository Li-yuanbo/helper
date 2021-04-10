package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"helper/model"
	"helper/store/mysql"
	"helper/utils"
)

func RegisterUser(req model.RegisterUserReq, c *gin.Context) (model.RegisterUserResp, error){
	var resp model.RegisterUserResp
	//加密
	req.Password = utils.MD5(req.Password)
	//写数据库
	user, err := mysql.AddUser(req, mysql.WriteDB())
	if err != nil {
		resp.Res = utils.NewRes(10003, utils.ERR_CODE[10003])
		resp.User = nil
		return resp, err
	}
	//写session
	session := sessions.Default(c)
	if session.Get("user") == nil{
		session.Set("user", user.Id)
		err := session.Save()
		if err != nil {
			resp.Res = utils.NewRes(10006, utils.ERR_CODE[10006])
			resp.User = nil
			return resp, err
		}
	}
	//构造返回参数
	resp.Res = utils.NewRes(10000, "SUCCESS")
	resUser := &model.UserInfoModel{
		Id:         user.Id,
		UserName:   user.UserName,
		Password:   user.Password,
		Name:       user.Name,
		Phone:      user.Phone,
		UserType:   user.UserType,
		Gender:     user.Gender,
		Age:        user.Age,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	}
	resp.User = resUser
	return resp, nil
}