package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

var err error

type Users struct {
	Username string
	Password string
	Avatar   string
	Phone    int
	Token    string
	Age      int
	Sex      int
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
	db.Model(&users).Where("id = ?", 4).Update("name", "hello")
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
