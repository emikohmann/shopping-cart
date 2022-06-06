package app

import (
	"github.com/emikohmann/shopping-cart/api/src/api/clients/database"
	"github.com/emikohmann/shopping-cart/api/src/api/config"
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
	"os"
	"time"
)

func Start() {
	router := gin.New()

	// configs
	dbConfig := config.DatabaseConfig{
		UserName:  os.Getenv("DATABASE_USER_NAME"),
		Password:  os.Getenv("DATABASE_PASSWORD"),
		Host:      os.Getenv("DATABASE_HOST"),
		Port:      os.Getenv("DATABASE_PORT"),
		Schema:    os.Getenv("DATABASE_SCHEMA"),
		Charset:   os.Getenv("DATABASE_CHARSET"),
		ParseTime: os.Getenv("DATABASE_PARSE_TIME"),
		Loc:       os.Getenv("DATABASE_LOC"),
	}
	authConfig := config.AuthConfig{
		TokenSigningKey:        os.Getenv("AUTH_TOKEN_SIGNING_KEY"),
		TokenExpirationSeconds: os.Getenv("AUTH_TOKEN_EXPIRATION_SECONDS"),
	}
	serverConfig := config.ServerConfig{
		Port: os.Getenv("SERVER_PORT"),
	}

	// clients
	dbSchemas := []interface{}{&users.User{}, &items.Item{}}
	dbClient := database.NewDBClient(dbConfig, &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	}, dbSchemas...)

	// services
	usersService := usersService.NewServiceImpl(dbClient)
	itemsService := itemsService.NewServiceImpl(dbClient)
	authService := authService.NewServiceJWT(usersService, authConfig)

	// controllers
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

	if err := router.Run(serverConfig.Port); err != nil {
		logger.Panic("Error running application", err)
	}
}
