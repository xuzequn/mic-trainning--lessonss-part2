package internal

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"mic-trainning-lessons-part2/model"
	"os"
	"time"
)

var DB *gorm.DB
var err error

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbName"`
	UserName string `mapstructure:"userName"`
	Password string `mapstructure:"password"`
}

func InitDB() {
	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //io.writer
		logger.Config{
			SlowThreshold:             time.Second, //slowsql
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        //忽略ErrorRecordNotFound(记录未找到)报错
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	host := AppConf.DBConfig.Host
	port := AppConf.DBConfig.Port
	name := AppConf.DBConfig.UserName
	password := AppConf.DBConfig.Password
	dbname := AppConf.DBConfig.DBName
	conn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", name, password, host, port, dbname)
	zap.S().Infof(conn)
	DB, err = gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newlogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表使用英文单数形式
		},
	})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	err = DB.AutoMigrate(&model.Category{}, &model.Brand{}, &model.Advertise{}, &model.Product{})
	if err != nil {
		fmt.Println(err)
	}
	zap.S().Info("已连接mysql")
}
