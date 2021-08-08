package models

import "github.com/jinzhu/gorm"


// 属性必须大写
type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(100)"`	// 默认即为username
	Password string
	Age      int
	RoleId   uint							// 列名会转换为role_id
	role     Role							// 结构体不会存入数据库
}