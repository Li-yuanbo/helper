package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
		resp.Res = utils.NewRes(10003, utils.ERR_CODE[10003])
		return &resp, err
	}
	//写session
	session := sessions.Default(c)
	if session.Get("org") == nil || session.Get("org") != organization.Id {
		session.Set("org", organization.Id)
		err := session.Save()
		if err != nil {
			log.Println("[service] save session err: ", err)
			resp.Res = utils.NewRes(10006, utils.ERR_CODE[10006])
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
