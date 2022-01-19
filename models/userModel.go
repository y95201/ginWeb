package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

var err error

type Users struct {
	Username string
	Password string
}

func GetProjects(c *gin.Context) (users []Users) {
	db := InitDB()
	db.AutoMigrate(&Users{})
	db.Where("id = ?", "3").First(&users)
	return users
}
func GetList(c *gin.Context) (users []Users) {
	db := InitDB()
	db.AutoMigrate(&Users{})
	db.Find(&users)
	return users
}
func GetUpdate(c *gin.Context) (users []Users) {
	db := InitDB()
	db.AutoMigrate(&Users{})
	db.Model(&users).Where("id = ?", 3).Update("name", "hello")
	return users
}
func GetDel(c *gin.Context) (users []Users) {
	db := InitDB()
	db.AutoMigrate(&Users{})
	db.Where("id = ?", 3).Delete(&users)
	return users
}
func GetAdds(c *gin.Context) (users []Users) {
	db := InitDB()
	db.AutoMigrate(&Users{})
	db.Model(&users).Where("id = ?", 3).Update("name", "hello")
	return users
}

func InitDB() *gorm.DB {

	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/think?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	db.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	db.LogMode(false)      //打印sql语句
	//开启连接池
	db.DB().SetMaxIdleConns(10)    //最大空闲连接
	db.DB().SetMaxOpenConns(100)   //最大连接数
	db.DB().SetConnMaxLifetime(30) //最大生存时间(s)
	return db
}

// func Logger() *logrus.Logger {
// 	now := time.Now()
// 	logFilePath := ""
// 	if dir, err := os.Getwd(); err == nil {
// 		logFilePath = dir + "/logs/"
// 	}
// 	if err := os.MkdirAll(logFilePath, 0777); err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	logFileName := now.Format("2006-01-02") + ".log"
// 	//日志文件
// 	fileName := path.Join(logFilePath, logFileName)
// 	if _, err := os.Stat(fileName); err != nil {
// 		if _, err := os.Create(fileName); err != nil {
// 			fmt.Println(err.Error())
// 		}
// 	}
// 	//写入文件
// 	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
// 	if err != nil {
// 		fmt.Println("err", err)
// 	}

// 	//实例化
// 	logger := logrus.New()

// 	//设置输出
// 	logger.Out = src

// 	//设置日志级别
// 	logger.SetLevel(logrus.DebugLevel)

// 	//设置日志格式
// 	logger.SetFormatter(&logrus.TextFormatter{
// 		TimestampFormat: "2006-01-02 15:04:05",
// 	})
// 	return logger
// }
