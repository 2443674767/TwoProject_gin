package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
	Sex  bool
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/ginclass?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Person{})

	//db.Create(&Person{
	//	Name: "xxz",
	//	Age:  111,
	//	Sex:  false,
	//})

	var person Person
	db.First(&person, "name = ?", "xxz")
	fmt.Println(person)

	//切片
	var person2 []Person
	db.Find(&person2, "age < ?", "30")
	fmt.Println(person2)

	//改
	//db.Where("id = ?", 1).First(&Person{}).Updates(Person{
	//	Name: "yy",
	//	Age:  2,
	//})
	//map 必须保证全部都有
	//db.Where("id = ?", 1).First(&Person{}).Updates(map[string]interface{}{
	//批量覆盖
	//db.Where("id in (?)", []int{1, 2}).Find(&[]Person{}).Updates(map[string]interface{}{
	//	"Name": "yy",
	//	"Age":  20,
	//	"Sex":  true,
	//})

	//删除
	//db.Delete(&Person{}, "id = ?", 1)	//软删除
	//db.Where("id in (?)", []int{1, 2}).Unscoped().Delete(&Person{})	//物理地址删除 无法查询到 防止脏数据过多，可以移表等

	//defer db.Close()

}
