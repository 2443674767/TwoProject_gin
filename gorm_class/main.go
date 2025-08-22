package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//type User struct {
//	gorm.Model
//	Name string `gorm:"primaryKey;column:user_name;type:varchar(100);not null"`
//}
//
//func (u User) TableName() string {
//	if u.Name == "aaa" { //特定名称返回表名
//		return "users"
//	} else {
//		return "gm_users"
//	}
//}

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

type Student struct {
	gorm.Model
	ClassID     uint
	IDCard      IDCard
	StudentName string
	Teachers    []Teacher `gorm:"many2many:student_classes;"` //关联表 student_classes
	//TeacherID uint
}

type IDCard struct {
	gorm.Model
	StudentID uint
	Num       int
}

type Teacher struct {
	gorm.Model
	TeacherName string
	//StudentID uint
	Students []Student `gorm:"many2many:student_classes;"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ginclass?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//db.AutoMigrate(&User{})

	db.AutoMigrate(&Class{}, &Student{}, &IDCard{}, &Teacher{})

	r := gin.Default()

	r.POST("student", func(c *gin.Context) {
		var student Student
		_ = c.BindJSON(&student)
		db.Create(&student)
	})

	r.GET("student/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var student Student
		_ = c.BindJSON(&student)
		//db.First(&student, "id = ?", id)

		//预加载
		db.Preload("Teachers").Preload("IDCard").First(&student, "id = ?", id)
		c.JSON(200, gin.H{
			"student": student,
		})
	})

	r.GET("/Class/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var class Class
		//db.First(&class, "id = ?", id)
		db.Preload("Students").Preload("Students.Teachers").First(&class, "id = ?", id)
		c.JSON(200, gin.H{
			"class": class,
		})

	})

	r.Run(":14005")

	//i := IDCard{
	//	Num: 123,
	//}
	//
	//s := Student{
	//	StudentName: "张hi就",
	//	IDCard:      i,
	//}
	//
	//t := Teacher{
	//	TeacherName: "王老师",
	//	//Students:    []Student{s},
	//}
	//
	//c := Class{
	//	ClassName: "开发班",
	//	Students:  []Student{s},
	//}
	//
	//_ = db.Create(&c).Error
	//
	//_ = db.Create(&t).Error

}
