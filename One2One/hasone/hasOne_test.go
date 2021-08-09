package hasone

import "gorm.io/gorm"

// 定义一个玩家结构
type Player struct {
	gorm.Model
	UID 		uint 		`gorm:"unique;comment:玩家唯一id"`
	Username 	string 		`gorm:"comment:玩家用户名"`
	Password	string		`gorm:"comment:密码"`
}


// 定义一个游戏角色
type Role struct {
	gorm.Model
	Name 	string 		`gorm:"comment:游戏角色名"`
	Level 	string		`gorm:"comment:游戏角色登记"`
	Coins  	uint32		`gorm:"comment:游戏角色金币数"`
}