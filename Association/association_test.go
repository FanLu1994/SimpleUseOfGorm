package Association

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
	err := db.AutoMigrate(&Player{},&Role{})		// 自动创建多对多中间表
	if err!=nil{
		fmt.Println("数据表迁移失败")
	}else{
		fmt.Println("数据表迁移成功")
	}
}


// 插入数据
func TestInsert(t *testing.T) {
	// 插入两个数据
	player1 := Player{
		Name: "玩家一号",
		Role: Role{
			Name: "马里奥",
		},
	}
	player2 := Player{
		Name: "玩家二号",
		Role: Role{
			Name: "路易吉",
		},
	}
	db.Create([]*Player{&player1,&player2})
}


// 测试关联查询
func TestFind(t *testing.T) {
	player := Player{}
	db.First(&player,1)

	// 查找所有匹配的关联记录
	role := Role{}
	_ = db.Model(&player).Association("Role").Find(&role)
	fmt.Println(role)

	// 替换关联
	newRole := Role{
		Name: "奇诺比奥",
	}
	db.Create(&newRole)
	err := db.Model(&player).Association("Role").Replace(&newRole, &role)
	fmt.Println(err)
}