package app

import (
	"github.com/emikohmann/shopping-cart/api/src/api/clients/database"
	itemsController "github.com/emikohmann/shopping-cart/api/src/api/controllers/items"
	usersController "github.com/emikohmann/shopping-cart/api/src/api/controllers/users"
	"github.com/emikohmann/shopping-cart/api/src/api/domain/items"
	"github.com/emikohmann/shopping-cart/api/src/api/domain/users"
	itemsService "github.com/emikohmann/shopping-cart/api/src/api/services/items"
	usersService "github.com/emikohmann/shopping-cart/api/src/api/services/users"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

func Start() {
	router := gin.New()

	// dependencies
	var (
		dbSchemas = []interface{}{&users.User{}, &items.Item{}}
		dbClient  = database.NewDBClient("root:@tcp(127.0.0.1:3306)/cart?charset=utf8mb4&parseTime=True&loc=Local", &gorm.Config{
			NowFunc: time.Now().UTC,
			Logger:  gormLogger.Default.LogMode(gormLogger.Info),
		}, dbSchemas...)
		usersService    = usersService.NewServiceImpl(dbClient)
		usersController = usersController.NewControllerImpl(usersService)
		itemsService    = itemsService.NewServiceImpl(dbClient)
		itemsController = itemsController.NewControllerImpl(itemsService)
	)

	// url mappings
	router.POST("/users", usersController.Create)
	router.POST("/login", usersController.Login)
	router.POST("/items", itemsController.Create)
	router.GET("/items/:id", itemsController.Get)
	router.GET("/items/search", itemsController.Search)

	if err := router.Run(":8080"); err != nil {
		logger.Panic("Error running application", err)
	}
}
