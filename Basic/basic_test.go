package Basic

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)


// 测试连接mysql
func TestConnectMysql(t *testing.T) {
	dsn := "root:1234@(127.0.0.1:3306)/SimpleGorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil{
		panic(err)
	}

	exist := db.Migrator().HasTable(&Player{})  // 查看Player表是否存在
	fmt.Println(exist)
}


// 测试创建数据表
func TestCreateTable(t *testing.T){
	dsn := "root:1234@(127.0.0.1:3306)/SimpleGorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil{
		panic(err)
	}

	db.Create(&Player{})  		// 创建player表
	exist := db.Migrator().HasTable(&Player{})  // 查看Player表是否存在
	fmt.Println(exist)
}


// 测试数据表迁移
func TestAutoMigrate(t *testing.T){
	//db, err := gorm.Open("mysql", "root:1234@(127.0.0.1:3306)/SimpleGorm?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := "root:1234@tcp(127.0.0.1:3306)/SimpleGorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err!= nil{
		panic(err)
	}

	if db.Migrator().HasTable(&Player{}){
		_ = db.Migrator().DropTable(&Player{}) // 删除表
	}

	err = db.AutoMigrate(&Player{})              // 数据库迁移\支持多个表同时
	if err!=nil{
		fmt.Println(err)
	}
	exist := db.Migrator().HasTable(&Player{}) // 查看Player表是否存在
	fmt.Println(exist)
	if exist{
		fmt.Println("players 表创建成功")
	}else{
		fmt.Println("players 表创建失败")
	}
}

var(
	dsn = "root:1234@(127.0.0.1:3306)/SimpleGorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
)


// 测试插入数据
func TestInsert(t *testing.T){
	player := Player{
		Username: "夏目",
		Password: "1234",
		UID: 12345678,
	}

	result := db.Save(&player)
	if result.RowsAffected ==1{
		fmt.Println("插入成功")
	}else{
		fmt.Println("插入失败")
	}

}

// 测试删除数据(软删除,注意deleteat字段)
func TestDelete(t *testing.T){
	player := Player{}
	db.First(&player)		// 删除第一条数据

	result := db.Delete(&player)
	if result.RowsAffected ==1{
		fmt.Println("删除成功")
	}else{
		fmt.Println("删除失败")
	}
}


// 更新数据
func TestUpdate(t *testing.T){
	player := Player{}
	db.First(&player)		// 会自动忽略删除的数据

	player.Username = "夏目修改名字"
	result := db.Updates(&player)
	if result.RowsAffected ==1{
		fmt.Println("更新成功")		// 注意updateat字段也会自动更新
	}else{
		fmt.Println("更新失败")
	}
}


// 查询单条数据
func TestQueryOne(t *testing.T){
	player := Player{}
	// 查询首条记录
	db.First(&player)
	fmt.Println(player)

	// 主键查询
	newPlayer := Player{}
	db.First(&newPlayer,4)
	fmt.Println(newPlayer)

	//db.Where("Username = ?","xiamu").First(&player)
}


// 查询多条数据
func TestQueryMany(t *testing.T){
	// 查询所有数据
	var players []Player
	db.Find(&players)
	fmt.Println(players)

	// Where查询数据  模糊查询
	var newPlayers []Player
	db.Debug().Where("username like ?","%夏目%").Find(&newPlayers)  // 加入DeBug()方便打印sql进行调试
	fmt.Println(newPlayers)
}

// 还有很多查询方法,不再赘述