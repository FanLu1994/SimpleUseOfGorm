package Basic

import (
	"gorm.io/gorm"
	"time"
)

// 定义一个玩家结构
type Player struct {
	ID 			uint 		`gorm:"primaryKey;comment:主键"`
	UID 		uint 		`gorm:"unique;comment:玩家唯一id"`
	Username 	string 		`gorm:"comment:玩家用户名"`
	Password	string		`gorm:"comment:密码"`
	CreatedAt 	time.Time	`gorm:"comment:创建时间;autoCreateTime"`
	UpdateAt 	time.Time	`gorm:"comment:创建时间;autoUpdateTime"`
	DeleteAt 	gorm.DeletedAt	`gorm:"删除时间"`
}
