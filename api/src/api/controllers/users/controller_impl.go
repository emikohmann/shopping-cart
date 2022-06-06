package users

import (
	domain "github.com/emikohmann/shopping-cart/api/src/api/domain/users"
	service "github.com/emikohmann/shopping-cart/api/src/api/services/users"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/apierrors"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controllerImpl struct {
	usersService service.Service
}

func NewControllerImpl(usersService service.Service) controllerImpl {
	logger.Info("Users controller successfully initialized")
	return controllerImpl{
		usersService: usersService,
	}
}

func (c controllerImpl) Create(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		logger.Error("Invalid user", err)
		apiErr := apierrors.NewBadRequestAPIError("invalid user")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	user, apiErr := c.usersService.Create(user)
	if apiErr != nil {
		logger.Error("Error creating user", apiErr)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	logger.Info("User successfully created")
	ctx.JSON(http.StatusCreated, user)
}
