package mysql

import (
	"github.com/jinzhu/gorm"
	"helper/model"
	"log"
	"time"
)

type Organization struct {
	Id                int64  `json:"id"`
	OrgName           string `json:"org_name"`      //机构组织名称
	OrgPic            string `json:"org_pic"`       //机构图片
	OrgUserName       string `json:"org_user_name"` //机构用户名
	OrgPassword       string `json:"org_password"`
	OrgRepresentative string `json:"org_representative"` //机构代表人名称
	OrgPhone          string `json:"org_phone"`          //机构代表人联系方式
	OrgProvince       string `json:"org_province"`
	OrgTown           string `json:"org_town"`
	OrgCountry        string `json:"org_country"`
	OrgAddress        string `json:"org_address"`
	OrgDesc           string `json:"org_desc"`
	CreateTime        int64  `json:"create_time"`
	UpdateTime        int64  `json:"update_time"`
}

func (*Organization) TableName() string {
	return "organization"
}

func AddOrganization(req *model.RegisterOrganizationReq, db *gorm.DB) (*Organization, error) {
	orgModel := Organization{
		OrgName:           req.OrgName,
		OrgPic:            req.OrgPic,
		OrgUserName:       req.OrgUserName,
		OrgPassword:       req.OrgPassword,
		OrgRepresentative: req.OrgRepresentative,
		OrgPhone:          req.OrgPhone,
		OrgProvince:       req.OrgProvince,
		OrgTown:           req.OrgTown,
		OrgCountry:        req.OrgCountry,
		OrgAddress:        req.OrgAddress,
		OrgDesc:           req.OrgDesc,
		CreateTime:        time.Now().Unix(),
		UpdateTime:        time.Now().Unix(),
	}
	if err := db.Create(&orgModel).Error; err != nil {
		log.Println("[db] add org err: ", err, ". org: ", orgModel)
		return nil, err
	}
	log.Println("[db] add org success")
	return &orgModel, nil
}
