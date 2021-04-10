package model

import "helper/utils"

type ActivityModel struct {
	Id                  int64  `json:"id"`
	OrgId               int64  `json:"org_id"`
	Title               string `json:"title"`
	Desc                string `json:"desc"`
	Representative      string `json:"representative"`
	RepresentativePhone string `json:"representative_phone"`
	TargetNum           int64  `json:"target_num"`
	Tag                 string `json:"tag"`
	Place               string `json:"place"`
	EndTime             int64  `json:"end_time"`
	EndStatus           int64  `json:"end_status"`
	CreateTime          int64  `json:"create_time"`
	UpdateTime          int64  `json:"update_time"`
}

type PublishActivityReq struct {
	OrgId               int64  `json:"org_id"`
	Title               string `json:"title"`
	Desc                string `json:"desc"`
	Representative      string `json:"representative"`
	RepresentativePhone string `json:"representative_phone"`
	TargetNum           int64  `json:"target_num"`
	Tag                 string `json:"tag"`
	Place               string `json:"place"`
	EndTime             int64  `json:"end_time"`
}

type PublishActivityResp struct {
	ResStatus *utils.Res     `json:"res_status"`
	Activity  *ActivityModel `json:"activity"`
}
