package model

import "helper/utils"

type OrganizationModel struct {
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

type RegisterOrganizationReq struct {
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
}

type RegisterOrganizationResp struct {
	Res          *utils.Res         `json:"res"`
	Organization *OrganizationModel `json:"organization"`
}
