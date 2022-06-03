package items

import (
	"errors"
	"fmt"
	"github.com/emikohmann/shopping-cart/api/src/api/domain/items"
	"github.com/emikohmann/shopping-cart/api/src/api/domain/search"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/apierrors"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/logger"
	"gorm.io/gorm"
	"strings"
)

type serviceImpl struct {
	dbClient *gorm.DB
}

func NewServiceImpl(dbClient *gorm.DB) serviceImpl {
	logger.Info("Items service successfully initialized")
	return serviceImpl{
		dbClient: dbClient,
	}
}

func (s serviceImpl) Create(item items.Item) (items.Item, apierrors.APIError) {
	result := s.dbClient.Create(&item)
	if result.Error != nil {
		logger.Error("Error creating item", result.Error)
		return items.Item{}, apierrors.NewInternalServerAPIError("error creating item")
	}

	return item, nil
}

func (s serviceImpl) Get(id int64) (items.Item, apierrors.APIError) {
	var item items.Item
	result := s.dbClient.First(&item, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return items.Item{}, apierrors.NewNotFoundAPIError("item not found")
	}

	if result.Error != nil {
		logger.Error("Error getting item", result.Error)
		return items.Item{}, apierrors.NewInternalServerAPIError("error getting item")
	}

	return item, nil
}

func (s serviceImpl) Search(query search.Query) (search.Result, apierrors.APIError) {
	var finalQuery *gorm.DB
	f := 0
	for field, values := range query.Data {
		var fieldQuery *gorm.DB
		for i, value := range values {
			if i == 0 {
				fieldQuery = s.dbClient.Where(fmt.Sprintf("%s LIKE '%%%s%%'", field, value))
			} else {
				fieldQuery = fieldQuery.Or(fmt.Sprintf("%s LIKE '%%%s%%'", field, value))
			}
		}
		if fieldQuery == nil {
			continue
		}
		if f == 0 {
			finalQuery = s.dbClient.Where(fieldQuery)
		} else {
			finalQuery = finalQuery.Where(fieldQuery)
		}
		f++
	}

	var target *gorm.DB
	results := make([]items.Item, 0)
	if finalQuery != nil {
		target = finalQuery
	} else {
		target = s.dbClient
	}

	result := target.Limit(query.Limit).Offset(query.Offset).Find(&results)
	if result.Error != nil {
		logger.Error("Error searching items", result.Error)
		if strings.Contains(strings.ToLower(result.Error.Error()), "unknown column") {
			return search.Result{}, apierrors.NewBadRequestAPIError("invalid field name")
		}
		return search.Result{}, apierrors.NewInternalServerAPIError("error searching items")
	}

	return search.Result{
		Paging: search.Paging{
			Limit:  query.Limit,
			Offset: query.Offset,
		},
		Results: results,
	}, nil
}
