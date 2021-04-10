package mysql

type UserInfo struct {
	Id			int64	`gorm:"column:id"`
	UserName	string  `gorm:"column:user_name"`	//用户名
	Password	string  `gorm:"column:password"`
	Name		string  `gorm:"column:name"`		//用户昵称
	Phone		string  `gorm:"column:phone"`
	UserType	int64   `gorm:"column:user_type"`	//0-普通用户 1-admin用户
	Gender		int64   `gorm:"column:gender"`
	Age			int64   `gorm:"column:age"`
	CreateTime	int64   `gorm:"column:create_time"`
	UpdateTime	int64   `gorm:"column:update_time"`
}

func (*UserInfo)TableName() string {
	return "user_info"
}