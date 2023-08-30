package models

type User struct {
	Id         string `gorm:"unique"`
	Username   string `gorm:"unique;not null;type:varchar(100)"`
	Password   string `gorm:"unique;not null;type:varchar(100)"`
	Role       string `gorm:"column:role"`
	Updatetime int    `gorm:"column:update_time"`
	CreatedAt  int    `gorm:"column:created_at"`
}
