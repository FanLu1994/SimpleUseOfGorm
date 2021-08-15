package Many2Many

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

var(
	dsn = "root:1234@(127.0.0.1:3306)/SimpleGorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,	//禁用物理外键
	})
)

func TestAutoMigrate(t *testing.T) {
	err := db.AutoMigrate(&Player{},&Gamepad{})		// 自动创建多对多中间表
	if err!=nil{
		fmt.Println("数据表迁移失败")
	}else{
		fmt.Println("数据表迁移成功")
	}
}

func TestInsert(t *testing.T) {
	gamepad1 := Gamepad{
		Brand: "Nintendo",
	}

	gamepad2 := Gamepad{
		Brand: "PS4",
	}
	db.Create([]*Gamepad{&gamepad1,&gamepad2})		// 插入手柄数据

	gamepads := make([]Gamepad,0)
	db.Find(&gamepads)

	player1 := Player{
		Name: "马里奥",
		Gamepads: gamepads,		// 每个player都有各种手柄
	}

	player2 := Player{
		Name: "路易吉",
		Gamepads: gamepads,
	}

	result := db.Create([]*Player{&player1,&player2})  //插入数据,gorm自动更新中间表
	if result.RowsAffected==0 {
		fmt.Println("插入失败")
	}else{
		fmt.Println("插入成功",result.RowsAffected)
	}
}
