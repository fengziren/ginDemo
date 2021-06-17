package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	UserName string
	Name     string
	Password string
	Tel      string
	Describe string
}

func main() {
	dsn := "root:ROOT@(localhost)/gindemo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 自动迁移
	db.AutoMigrate(&User{})
	// u1 := User{gorm.Model{CreatedAt: time.Now()}, "玛卡巴卡", "张三", "123456", "0803-123456", "追风少年"}
	// u2 := User{gorm.Model{ID:2,CreatedAt:time.Now()}, "唔西迪西", "李四", "123456", "0803-124456", "飞飞飞"}
	// 创建记录
	// db.Create(&u1)
	// db.Create(&u2)
	// 查询
	var u = new(User)
	db.First(u)
	fmt.Println(u)

	var uu User
	db.Find(&uu, "Name=?", "李四")
	fmt.Printf("%#v\n", uu)

	// 更新
	// db.Model(&u).Update("Tel", "01111110101")
	// 删除
	// db.Delete(&u)
}
