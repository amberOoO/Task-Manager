package utils

import (
	"sync"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// viper相关配置
func LoadConfig() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// logrus.SetLevel(logrus.DebugLevel)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}


// 获取issue数据库连接
var load_issue_db_once sync.Once
var issue_db *gorm.DB

func GetIssueDBConnection() *gorm.DB{
	load_issue_db_once.Do(func() {
		meta := viper.GetStringMapString("postgres")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			meta["host"],
			meta["user"],
			meta["pwd"],
			meta["db"],
			meta["port"])
		fmt.Println(dsn)
		var err error
		if issue_db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		}); err != nil {
			panic(err)
		}
	})
	return issue_db
}
