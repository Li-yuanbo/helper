package mysql

import (
	"github.com/jinzhu/gorm"
	"helper/model"
	"log"
	"time"
)

type Activity struct {
	Id                  int64  `gorm:"column:id"`
	OrgId               int64  `grom:"column:org_id"`
	Title               string `gorm:"column:title"`
	Desc                string `gorm:"column:desc"`
	Representative      string `gorm:"column:representative"`
	RepresentativePhone string `gorm:"column:representative_phone"`
	TargetNum           int64  `gorm:"column:target_num"`
	Tag                 string `gorm:"column:tag"`
	Place               string `gorm:"column:place"`
	EndTime             int64  `gorm:"column:end_time"`
	EndStatus           int64  `gorm:"column:end_status"`
	CreateTime          int64  `gorm:"column:create_time"`
	UpdateTime          int64  `gorm:"column:update_time"`
}

func (*Activity) TableName() string {
	return "activity"
}

func AddActivity(req *model.PublishActivityReq, db *gorm.DB) (*Activity, error) {
	activityModel := Activity{
		OrgId:               req.OrgId,
		Title:               req.Title,
		Desc:                req.Title,
		Representative:      req.Representative,
		RepresentativePhone: req.RepresentativePhone,
		TargetNum:           req.TargetNum,
		Tag:                 req.Tag,
		Place:               req.Place,
		EndTime:             req.EndTime,
		EndStatus:           0,
		CreateTime:          time.Now().Unix(),
		UpdateTime:          time.Now().Unix(),
	}
	if err := db.Create(&activityModel).Error; err != nil {
		log.Println("[db] publish activity err: ", err)
		return nil, err
	}
	log.Println("[db] publish activity success")
	return &activityModel, nil
}
