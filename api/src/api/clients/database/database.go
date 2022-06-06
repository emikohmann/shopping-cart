package database

import (
	"fmt"
	"github.com/emikohmann/shopping-cart/api/src/api/config"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	connectionDSN = "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s"
)

func NewDBClient(dbConfig config.DatabaseConfig, config *gorm.Config, schemas ...interface{}) *gorm.DB {
	dsn := fmt.Sprintf(connectionDSN,
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Schema,
		dbConfig.Charset,
		dbConfig.ParseTime,
		dbConfig.Loc)
	client, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		logger.Panic("Error creating MySQL client using GORM", err)
	}
	if err := client.AutoMigrate(schemas...); err != nil {
		logger.Panic("Error migrating schemas", err)
	}
	logger.Info("MySQL database successfully initialized")
	return client
}
