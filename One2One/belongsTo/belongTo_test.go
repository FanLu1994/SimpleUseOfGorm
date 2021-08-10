package belongsTo

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


// hasOne数据表迁移测试，默认外键,指定外键,自定义引用等等
func TestAutoMigrate(t *testing.T){
	err := db.AutoMigrate(&Player{})		//会自动创建包含的结构体表
	if err!=nil{
		fmt.Println("数据表迁移失败")
	}else{
		fmt.Println("数据表迁移成功")
	}
}


func TestInsert(t *testing.T){
	player1 := Player{
		Username: "周芷若",
		Password: "1234",
		Sword: Sword{
			Name: "倚天剑",
		},
	}

	player2 := Player{
		Username: "张无忌",
		Password: "1234",
		Sword: Sword{
			Name: "屠龙刀",
		},
	}

	db.Create([]Player{player1,player2})
}


func TestQuery(t *testing.T){
	players := []Player{}

	db.Debug().Preload("Sword").Find(&players)

	for _,player := range players{
		fmt.Println(player)
	}
}