package app

import (
	"github.com/emikohmann/shopping-cart/api/src/api/clients/database"
	authController "github.com/emikohmann/shopping-cart/api/src/api/controllers/auth"
	itemsController "github.com/emikohmann/shopping-cart/api/src/api/controllers/items"
	usersController "github.com/emikohmann/shopping-cart/api/src/api/controllers/users"
	"github.com/emikohmann/shopping-cart/api/src/api/domain/items"
	"github.com/emikohmann/shopping-cart/api/src/api/domain/users"
	authService "github.com/emikohmann/shopping-cart/api/src/api/services/auth"
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
	dbSchemas := []interface{}{&users.User{}, &items.Item{}}
	dbClient := database.NewDBClient("root:@tcp(127.0.0.1:3306)/cart?charset=utf8mb4&parseTime=True&loc=Local", &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	}, dbSchemas...)

	usersService := usersService.NewServiceImpl(dbClient)
	itemsService := itemsService.NewServiceImpl(dbClient)
	authService := authService.NewServiceJWT(usersService, []byte("4h0pp1ngk4r7"))

	usersController := usersController.NewControllerImpl(usersService)
	itemsController := itemsController.NewControllerImpl(itemsService)
	authController := authController.NewControllerImpl(authService)

	// url mappings
	router.POST("/users", usersController.Create)
	router.POST("/auth/login", authController.Login)
	router.POST("/auth/validate", authController.Validate)
	router.POST("/items", itemsController.Create)
	router.GET("/items/:id", itemsController.Get)
	router.GET("/items/search", itemsController.Search)

	if err := router.Run(":8080"); err != nil {
		logger.Panic("Error running application", err)
	}
}
