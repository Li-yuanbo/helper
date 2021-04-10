package service

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"helper/model"
	"helper/store/mysql"
	"helper/utils"
)

func RegisterUser(req model.RegisterUserReq, c *gin.Context) (model.RegisterUserResp, error) {
	var resp model.RegisterUserResp
	//加密
	req.Password = utils.MD5(req.Password)
	//写数据库
	user, err := mysql.AddUser(req, mysql.WriteDB())
	if err != nil {
		resp.ResStatus = utils.NewRes(10003, utils.ERR_CODE[10003])
		resp.User = nil
		return resp, err
	}
	//写session
	session := sessions.Default(c)
	if session.Get("user") == nil || session.Get("user").(int64) != user.Id {
		session.Set("user", user.Id)
		err := session.Save()
		if err != nil {
			resp.ResStatus = utils.NewRes(10006, utils.ERR_CODE[10006])
			resp.User = nil
			return resp, err
		}
	} else {
		resp.ResStatus = utils.NewRes(20003, utils.ERR_CODE[20003])
		return resp, err
	}
	//构造返回参数
	resp.ResStatus = utils.NewRes(10000, "SUCCESS")
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

func GetUserInfos(c *gin.Context, req *model.GetUserInfosReq) (*model.GetUserInfosResp, error) {
	var resp model.GetUserInfosResp
	tx := mysql.WriteDB().Begin()
	defer tx.RollbackUnlessCommitted()
	//获取用户总数
	totalNum, err := mysql.GetUserCount(tx)
	if err != nil {
		resp.ResStatus = utils.NewRes(10003, utils.ERR_CODE[10003])
		return &resp, err
	}
	resp.TotalUserNum = totalNum
	//计算页数情况
	if totalNum%req.Limit == 0 {
		resp.TotalPage = totalNum / req.Limit
	} else {
		resp.TotalPage = totalNum/req.Limit + 1
	}
	resp.CurPage = req.CurPage + 1
	//分页获取用户
	users, err := mysql.GetUsersByPage(req, tx)
	if err != nil {
		resp.ResStatus = utils.NewRes(10003, utils.ERR_CODE[10003])
		return &resp, err
	}
	for _, user := range users {
		userModel := model.UserInfoModel{
			Id:         user.Id,
			UserName:   user.UserName,
			Password:   user.Password,
			Name:       user.Name,
			Phone:      user.Phone,
			UserType:   user.UserType,
			Gender:     user.Gender,
			Age:        user.Age,
			CreateTime: user.CreateTime,
			UpdateTime: user.UserType,
		}
		resp.Users = append(resp.Users, &userModel)
	}
	resp.ResStatus = utils.NewRes(10000, "SUCCESS")
	return &resp, nil
}

func LoginUser(c *gin.Context, req *model.LoginUserReq) (*model.LoginUserResp, error) {
	var resp model.LoginUserResp
	user, err := mysql.GetUserByName(req, mysql.WriteDB())
	//通过user_name未找到用户信息
	if err != nil && gorm.IsRecordNotFoundError(err) {
		resp.ResStatus = utils.NewRes(20002, utils.ERR_CODE[20002])
		return &resp, err
	} else if err != nil {
		resp.ResStatus = utils.NewRes(10003, utils.ERR_CODE[10003])
		return &resp, err
	}
	//密码错误
	if user.Password != utils.MD5(req.Password) {
		resp.ResStatus = utils.NewRes(20001, utils.ERR_CODE[20001])
		return &resp, errors.New("password_error")
	}
	//判断是否已经登录
	session := sessions.Default(c)
	if session.Get("user") == nil || session.Get("user").(int64) != user.Id {
		session.Set("user", user.Id)
		err := session.Save()
		if err != nil {
			resp.ResStatus = utils.NewRes(10006, utils.ERR_CODE[10006])
			resp.User = nil
			return &resp, err
		}
	} else {
		resp.ResStatus = utils.NewRes(20003, utils.ERR_CODE[20003])
		return &resp, errors.New("user_already_login")
	}
	userModel := model.UserInfoModel{
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
	resp.ResStatus = utils.NewRes(10000, "SUCCESS")
	resp.User = &userModel
	return &resp, nil
}

func UnLoginUser(c *gin.Context, req *model.UnLoginUserReq) (*model.UnLoginUserResp, error) {
	var resp model.UnLoginUserResp
	session := sessions.Default(c)
	if session.Get("user") == nil || session.Get("user").(int64) != req.Id {
		resp.ResStatus = utils.NewRes(20004, utils.ERR_CODE[20004])
		return &resp, errors.New("get session err")
	} else if session.Get("user").(int64) == req.Id {
		//删除session
		session.Delete("user")
		err := session.Save()
		if err != nil {
			resp.ResStatus = utils.NewRes(10006, utils.ERR_CODE[10006])
			return &resp, err
		}
	}
	resp.ResStatus = utils.NewRes(10000, "SUCCESS")
	return &resp, nil
}
