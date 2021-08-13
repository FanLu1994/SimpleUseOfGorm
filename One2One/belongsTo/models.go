package belongsTo

import "gorm.io/gorm"

type Player struct {
	gorm.Model				// gorm提供的结构体,自带id\Createat\updateat\deleteat
	Username 	string 		`gorm:"comment:玩家用户名"`
	Password	string		`gorm:"comment:密码"`

	// 1.最普通
	SwordID		uint		`gorm:"comment:宝剑ID"`
	Sword		Sword

	// 2.自定义外键名字
	//MySwordID	string
	//Sword		Sword		`gorm:"foreignKey:MySwordID"`		//修改外键的名字

	// 3.自定义外键引用
	//SwordID		uint		`gorm:"comment:宝剑名字"`
	//Sword     Sword	 `gorm:"references:Name"`	// SwordID保存的即变为Sword.Name
}

type Sword struct {
	gorm.Model
	Name     	string		`gorm:"comment:宝剑名;unique"`
}