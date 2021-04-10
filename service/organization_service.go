package service

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"helper/model"
	"helper/store/mysql"
	"helper/utils"
	"log"
)

func RegisterOrganization(c *gin.Context, req *model.RegisterOrganizationReq) (*model.RegisterOrganizationResp, error) {
	var resp model.RegisterOrganizationResp
	//密码加密
	req.OrgPassword = utils.MD5(req.OrgPassword)
	//写数据库
	organization, err := mysql.AddOrganization(req, mysql.WriteDB())
	if err != nil {
		log.Println("[service] register org err: ", err)
		resp.ResStatus = utils.NewRes(10003, utils.ERR_CODE[10003])
		return &resp, err
	}
	//写session
	session := sessions.Default(c)
	if session.Get("org") == nil || session.Get("org") != organization.Id {
		session.Set("org", organization.Id)
		err := session.Save()
		if err != nil {
			log.Println("[service] save session err: ", err)
			resp.ResStatus = utils.NewRes(10006, utils.ERR_CODE[10006])
			return &resp, err
		}
	}
	//封装返回
	orgModel := model.OrganizationModel{
		Id:                organization.Id,
		OrgName:           organization.OrgName,
		OrgPic:            organization.OrgPic,
		OrgUserName:       organization.OrgUserName,
		OrgPassword:       organization.OrgPassword,
		OrgRepresentative: organization.OrgRepresentative,
		OrgPhone:          organization.OrgPhone,
		OrgProvince:       organization.OrgProvince,
		OrgTown:           organization.OrgTown,
		OrgCountry:        organization.OrgCountry,
		OrgAddress:        organization.OrgAddress,
		OrgDesc:           organization.OrgDesc,
		CreateTime:        organization.CreateTime,
		UpdateTime:        organization.UpdateTime,
	}
	resp.Organization = &orgModel
	return &resp, nil
}

func LoginOrg(c *gin.Context, req *model.LoginOrgReq) (*model.LoginOrgResp, error) {
	var resp model.LoginOrgResp
	organization, err := mysql.GetOrgByName(req, mysql.WriteDB())
	if err != nil && gorm.IsRecordNotFoundError(err) {
		resp.ResStatus = utils.NewRes(30001, utils.ERR_CODE[30001])
		return &resp, err
	} else if err != nil {
		resp.ResStatus = utils.NewRes(10003, utils.ERR_CODE[10003])
		return &resp, err
	}
	//判断密码
	if organization.OrgPassword != utils.MD5(req.OrgPassword) {
		resp.ResStatus = utils.NewRes(30002, utils.ERR_CODE[30002])
		return &resp, errors.New("password_error")
	}
	//判断是否已经登录
	session := sessions.Default(c)
	if session.Get("org") == nil || session.Get("org").(int64) != organization.Id {
		session.Set("org", organization.Id)
		err := session.Save()
		if err != nil {
			resp.ResStatus = utils.NewRes(10006, utils.ERR_CODE[10006])
			return &resp, err
		}
	} else {
		resp.ResStatus = utils.NewRes(30003, utils.ERR_CODE[30003])
		return &resp, errors.New("org_already_login")
	}
	resp.Org = &model.OrganizationModel{
		Id:                organization.Id,
		OrgName:           organization.OrgName,
		OrgPic:            organization.OrgPic,
		OrgUserName:       organization.OrgUserName,
		OrgPassword:       organization.OrgPassword,
		OrgRepresentative: organization.OrgRepresentative,
		OrgPhone:          organization.OrgPhone,
		OrgProvince:       organization.OrgProvince,
		OrgTown:           organization.OrgTown,
		OrgCountry:        organization.OrgCountry,
		OrgAddress:        organization.OrgAddress,
		OrgDesc:           organization.OrgDesc,
		CreateTime:        organization.CreateTime,
		UpdateTime:        organization.UpdateTime,
	}
	resp.ResStatus = utils.NewRes(10000, "SUCCESS")
	return &resp, nil
}

func UnLoginOrg(c *gin.Context, req *model.UnLoginOrgReq) (*model.UnLoginOrgResp, error) {
	var resp model.UnLoginOrgResp
	session := sessions.Default(c)
	if session.Get("org") == nil || session.Get("org").(int64) != req.Id {
		resp.ResStatus = utils.NewRes(30004, utils.ERR_CODE[30004])
		return &resp, errors.New("get session err")
	} else if session.Get("org").(int64) == req.Id {
		//删除session
		session.Delete("org")
		err := session.Save()
		if err != nil {
			resp.ResStatus = utils.NewRes(10006, utils.ERR_CODE[10006])
			return &resp, err
		}
	}
	resp.ResStatus = utils.NewRes(10000, "SUCCESS")
	return &resp, nil
}
