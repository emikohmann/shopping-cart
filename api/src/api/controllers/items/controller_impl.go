package items

import (
	domain "github.com/emikohmann/shopping-cart/api/src/api/domain/items"
	service "github.com/emikohmann/shopping-cart/api/src/api/services/items"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/apierrors"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/logger"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/search"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type controllerImpl struct {
	itemsService service.Service
}

func NewControllerImpl(itemsService service.Service) controllerImpl {
	logger.Info("Items controller successfully initialized")
	return controllerImpl{
		itemsService: itemsService,
	}
}

func (c controllerImpl) Create(ctx *gin.Context) {
	var item domain.Item

	if err := ctx.ShouldBindJSON(&item); err != nil {
		logger.Error("Invalid item", err)
		apiErr := apierrors.NewBadRequestAPIError("invalid item")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	item, apiErr := c.itemsService.Create(item)
	if apiErr != nil {
		logger.Error("Error creating item", apiErr)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	logger.Info("Item successfully created")
	ctx.JSON(http.StatusCreated, item)
}

func (c controllerImpl) Get(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Invalid id", err)
		apiErr := apierrors.NewBadRequestAPIError("invalid id")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	item, apiErr := c.itemsService.Get(id)
	if apiErr != nil {
		logger.Error("Error getting item", apiErr)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	logger.Info("Item successfully retrieved")
	ctx.JSON(http.StatusOK, item)
}

func (c controllerImpl) Search(ctx *gin.Context) {
	result, apiErr := c.itemsService.Search(search.ParseQuery(ctx))
	if apiErr != nil {
		logger.Error("Error searching items", apiErr)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	logger.Info("Search successfully completed")
	ctx.JSON(http.StatusOK, result)
}
