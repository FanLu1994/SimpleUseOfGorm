package Association

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name 		string
	Role		Role

}


type Role struct {
	gorm.Model
	Name   		string
	PlayerID	uint
}