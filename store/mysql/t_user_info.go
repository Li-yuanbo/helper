package mysql

import (
	"github.com/jinzhu/gorm"
	"helper/model"
	"log"
	"time"
)

type UserInfo struct {
	Id         int64  `gorm:"column:id"`
	UserName   string `gorm:"column:user_name"` //用户名
	Password   string `gorm:"column:password"`
	Name       string `gorm:"column:name"` //用户昵称
	Phone      string `gorm:"column:phone"`
	UserType   int64  `gorm:"column:user_type"` //0-普通用户 1-admin用户
	Gender     int64  `gorm:"column:gender"`
	Age        int64  `gorm:"column:age"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
}

func (*UserInfo) TableName() string {
	return "user_info"
}

func AddUser(req model.RegisterUserReq, db *gorm.DB) (*UserInfo, error) {
	userModel := UserInfo{
		UserName:   req.UserName,
		Password:   req.Password,
		Name:       req.Name,
		Phone:      req.Phone,
		UserType:   req.UserType,
		Gender:     req.Gender,
		Age:        req.Age,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	if err := db.Create(&userModel).Error; err != nil {
		log.Println("[db] register user err: ", err, ". user: ", userModel)
		return nil, err
	}
	log.Println("[db] register user success")
	return &userModel, nil
}

func GetUserCount(db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Model(&UserInfo{}).Count(&count).Error; err != nil {
		log.Println("[db] get user count err: ", err)
		return -1, err
	}
	log.Println("[db] get user count success")
	return count, nil
}

func GetUsersByPage(req *model.GetUserInfosReq, db *gorm.DB) ([]*UserInfo, error) {
	res := make([]*UserInfo, 0, 0)
	if err := db.Model(&UserInfo{}).Limit(req.Limit).Offset(req.Offset).Find(&res).Error; err != nil {
		log.Println("[db] get users err: ", err)
		return nil, err
	}
	log.Println("[db] get users success")
	return res, nil
}
