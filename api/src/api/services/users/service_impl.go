package users

import (
	"errors"
	"github.com/emikohmann/shopping-cart/api/src/api/domain/users"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/apierrors"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/hashing"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/logger"
	"gorm.io/gorm"
)

type serviceImpl struct {
	dbClient *gorm.DB
}

func NewServiceImpl(dbClient *gorm.DB) serviceImpl {
	logger.Info("Users service successfully initialized")
	return serviceImpl{
		dbClient: dbClient,
	}
}

func (s serviceImpl) Create(user users.User) (users.User, apierrors.APIError) {
	previous := s.dbClient.First(&user, "user_name = ?", user.UserName)
	if previous.Error != nil {
		if !errors.Is(previous.Error, gorm.ErrRecordNotFound) {
			logger.Error("Error creating user", previous.Error)
			return users.User{}, apierrors.NewInternalServerAPIError("error creating user")
		}
		user.Password = hashing.MD5(user.Password)
		result := s.dbClient.Create(&user)
		if result.Error != nil {
			logger.Error("Error creating user", previous.Error)
			return users.User{}, apierrors.NewInternalServerAPIError("error creating user")
		}
		return user, nil
	}
	return users.User{}, apierrors.NewBadRequestAPIError("user already exists")
}

func (s serviceImpl) Check(user users.User) (users.User, apierrors.APIError) {
	var found users.User
	result := s.dbClient.First(&found, "user_name = ? AND password = ?", user.UserName, hashing.MD5(user.Password))
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return users.User{}, apierrors.NewUnauthorizedAPIError("invalid credentials")
	}
	if result.Error != nil {
		logger.Error("Error checking user", result.Error)
		return users.User{}, apierrors.NewInternalServerAPIError("error checking user")
	}
	return found, nil
}
