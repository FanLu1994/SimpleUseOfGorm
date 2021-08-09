package hasone

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
	err := db.AutoMigrate(&Player{},&Role{},&Phone{},&Face{})
	if err!=nil{
		fmt.Println("数据表迁移失败")
	}else{
		fmt.Println("数据表迁移成功")
	}
}


// 测试插入数据
func TestInsert(t *testing.T){
	player := Player{
		UID: 12345,
		Username: "xiamu",
		Password: "1234",
		Role: Role{
			Name: "夏目",
			Level: "12",
			Coins: 100,
		},
		Phone: 	Phone{
			Brand: "xiaomi",
			Number: 138888888,
		},
		Face:	Face{
			FaceData: "testtest",
		},
	}

	result := db.Save(&player)
	if result.RowsAffected==0 {
		fmt.Println("插入失败")
	}else{
		fmt.Println("插入成功",result.RowsAffected)
	}
}

// 删除关联；清空关联


// 测试更新数据


// 测试添加关联


// 测试关联查询数据