package service

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"helper/model"
	"helper/store/mysql"
	"helper/utils"
)

func PublishActivity(c *gin.Context, req *model.PublishActivityReq) (*model.PublishActivityResp, error) {
	var resp model.PublishActivityResp
	session := sessions.Default(c)
	if session.Get("org") == nil || session.Get("org").(int64) != req.OrgId {
		resp.ResStatus = utils.NewRes(30004, utils.ERR_CODE[30004])
		return &resp, errors.New("org not login")
	}
	req.OrgId = session.Get("org").(int64)
	activity, err := mysql.AddActivity(req, mysql.WriteDB())
	if err != nil {
		resp.ResStatus = utils.NewRes(10003, utils.ERR_CODE[10003])
		return &resp, err
	}
	resp.ResStatus = utils.NewRes(10000, "SUCCESS")
	resp.Activity = &model.ActivityModel{
		Id:                  activity.Id,
		OrgId:               activity.OrgId,
		Title:               activity.Title,
		Desc:                activity.Desc,
		Representative:      activity.Representative,
		RepresentativePhone: activity.RepresentativePhone,
		TargetNum:           activity.TargetNum,
		Tag:                 activity.Tag,
		Place:               activity.Place,
		CreateTime:          activity.CreateTime,
		UpdateTime:          activity.UpdateTime,
	}
	return &resp, nil
}
