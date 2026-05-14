package model

type User struct {
	BaseModel
	UserID       uint64  `gorm:"column:user_id"`
	Username     string  `gorm:"column:username"`
	Password     string  `gorm:"column:password"`
	Mobile       string  `gorm:"column:mobile"`
	Nikename     string  `gorm:"column:nikename"`
	RegisterType int     `gorm:"column:register_type"`
	Balance      float64 `gorm:"column:balance"`
	IsDeleted    int     `gorm:"column:is_deleted"`
}

// TableName 指定用户模型对应的数据表名。
func (User) TableName() string {
	return "user"
}
