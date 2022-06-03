package users

import (
	"github.com/emikohmann/shopping-cart/api/src/api/domain/users"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/apierrors"
)

type Service interface {
	Create(user users.User) (users.User, apierrors.APIError)
	Login(user users.User) (users.User, apierrors.APIError)
}
