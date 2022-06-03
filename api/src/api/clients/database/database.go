package database

import (
	"github.com/emikohmann/shopping-cart/api/src/api/utils/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBClient(dsn string, config *gorm.Config, schemas ...interface{}) *gorm.DB {
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
