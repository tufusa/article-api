package infrastructure

import (
	"fmt"

	"github.com/tufusa/article-api/backend/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnector struct {
	Conn *gorm.DB
}

func NewDBConnector() *DBConnector {
	conf := config.LoadConfig()
	dsn := generateDSN(*conf.DB)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return &DBConnector{
		Conn: conn,
	}
}

func generateDSN(conf config.DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		conf.Username,
		conf.Password,
		conf.Addr,
		conf.DBName,
	)
}
