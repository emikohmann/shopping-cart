package auth

import (
	"github.com/emikohmann/shopping-cart/api/src/api/domain/auth"
	"github.com/emikohmann/shopping-cart/api/src/api/domain/users"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/apierrors"
)

type Service interface {
	Login(user users.User) (users.User, apierrors.APIError)
	Validate(auth auth.Auth) (auth.Auth, apierrors.APIError)
}
