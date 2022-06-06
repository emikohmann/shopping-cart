package auth

import (
	domain "github.com/emikohmann/shopping-cart/api/src/api/domain/auth"
	"github.com/emikohmann/shopping-cart/api/src/api/domain/users"
	service "github.com/emikohmann/shopping-cart/api/src/api/services/auth"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/apierrors"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controllerImpl struct {
	authService service.Service
}

func NewControllerImpl(authService service.Service) controllerImpl {
	return controllerImpl{
		authService: authService,
	}
}

func (c controllerImpl) Login(ctx *gin.Context) {
	var user users.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		logger.Error("Invalid user", err)
		apiErr := apierrors.NewBadRequestAPIError("invalid user")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	auth, apiErr := c.authService.Login(user)
	if apiErr != nil {
		logger.Error("Error logging user", apiErr)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	logger.Info("User successfully logged in")
	ctx.JSON(http.StatusOK, auth)
}

func (c controllerImpl) Validate(ctx *gin.Context) {
	var auth domain.Auth

	if err := ctx.ShouldBindJSON(&auth); err != nil {
		logger.Error("Invalid claim", err)
		apiErr := apierrors.NewBadRequestAPIError("invalid claim")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	validated, apiErr := c.authService.Validate(auth)
	if apiErr != nil {
		logger.Error("Error validating user", apiErr)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	logger.Info("Auth successfully validated")
	ctx.JSON(http.StatusOK, validated)
}
