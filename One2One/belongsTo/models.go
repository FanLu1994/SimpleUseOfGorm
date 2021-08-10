package belongsTo

import "gorm.io/gorm"

type Player struct {
	gorm.Model				// gorm提供的结构体,自带id\Createat\updateat\deleteat
	Username 	string 		`gorm:"comment:玩家用户名"`
	Password	string		`gorm:"comment:密码"`
	SwordID		uint		`gorm:"comment:宝剑ID"`
	Sword		Sword
}

type Sword struct {
	gorm.Model
	Name     	string		`gorm:"comment:公会名"`
}