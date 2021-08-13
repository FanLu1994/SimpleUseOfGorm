package One2Many

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
	err := db.AutoMigrate(&Player{}, &Hero{})
	if err!=nil{
		fmt.Println("数据表迁移失败")
	}else{
		fmt.Println("数据表迁移成功")
	}
}


func TestInsert(t *testing.T){
	hero1 := Hero{
		Name: "卡莎",
		HP: 300,
	}
	hero2 := Hero{
		Name: "瑞兹",
		HP: 500,
	}

	player := Player{
		Username: "夏目",
		Password: "1234",
		Heros: []Hero{hero1,hero2},
	}

	result := db.Create(&player)
	fmt.Println("插入行数：",result.RowsAffected)
}