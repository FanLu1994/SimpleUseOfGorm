package hasone

import (
	"gorm.io/gorm"
)

//------------------------------1-----------------------------------------
// 定义一个玩家结构
type Player struct {
	gorm.Model				// gorm提供的结构体,自带id\Createat\updateat\deleteat
	UID 		uint 		`gorm:"unique;comment:玩家唯一id"`
	Username 	string 		`gorm:"comment:玩家用户名"`
	Password	string		`gorm:"comment:密码"`
	Role 		Role
	Phone		Phone		`gorm:"foreignKey:UID"`					// 自定义外键
	Face		Face		`gorm:"foreignKey:UID;references:uid"`	// 自定义引用

}
// 定义一个游戏角色 使用playerId作为外键
type Role struct {
	gorm.Model
	PlayerId 	uint		// 默认使用user.uid作为外键
	Name 		string 		`gorm:"comment:游戏角色名"`
	Level 		string		`gorm:"comment:游戏角色登记"`
	Coins  		uint32		`gorm:"comment:游戏角色金币数"`
}
// 定义手机结构体 使用uid作为外键
type Phone struct {
	gorm.Model
	Brand     	string		`gorm:"comment:手机型号"`
	Number    	int			`gorm:"comment:手机号码"`
	UID			int
}
// 定义捏脸数据  使用Username作为外键引用
type Face struct {
	gorm.Model
	FaceData 	string 		 `gorm:"comment:捏脸数据"`
	UID			uint
}