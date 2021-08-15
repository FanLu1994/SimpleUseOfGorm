package Many2Many

import "gorm.io/gorm"

// 每一个player有多个手柄
type Player struct {
	gorm.Model
	Name 		string
	Gamepads	[]Gamepad 		`gorm:"many2many:player_gamepad"`
	Friends		[]Player		`gorm:"many2many:player_friends"`		//自引用
}

// 游戏手柄
type Gamepad struct {
	gorm.Model
	Brand 		string
}