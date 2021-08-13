package One2Many

import "gorm.io/gorm"

type Player struct {
	gorm.Model				// gorm提供的结构体,自带id\Createat\updateat\deleteat
	Username 	string 		`gorm:"comment:玩家用户名"`
	Password	string		`gorm:"comment:密码"`
	//Heros 		[]Hero

	// 第二种
	//Heros 		[]Hero		`gorm:"foreignKey:PlayerKey"`	// 定义外键的名字

	// 第三种
	Heros		[]Hero		`gorm:"foreignKey:PlayerName;references:Username"`	// 自定义外键引用
}


// hero   player可以拥有多个hero
type Hero struct {
	gorm.Model
	Name		string		`gorm:"comment:英雄名字"`
	HP			int32		`gorm:"comment:血量"`
	//PlayerID	uint

	// 第二种
	//PlayerKey	uint

	// 第三种
	PlayerName 	string

}
