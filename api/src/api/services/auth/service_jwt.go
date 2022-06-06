package auth

import (
	"errors"
	"github.com/emikohmann/shopping-cart/api/src/api/config"
	domain "github.com/emikohmann/shopping-cart/api/src/api/domain/auth"
	"github.com/emikohmann/shopping-cart/api/src/api/domain/users"
	usersService "github.com/emikohmann/shopping-cart/api/src/api/services/users"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/apierrors"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/logger"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

const (
	claimUserName  = "user_name"
	claimExpiresAt = "expires_at"
)

type serviceJWT struct {
	usersService    usersService.Service
	tokenSigningKey []byte
	tokenExpiration time.Duration
}

func NewServiceJWT(usersService usersService.Service, config config.AuthConfig) serviceJWT {
	seconds, err := strconv.ParseInt(config.TokenExpirationSeconds, 10, 64)
	if err != nil {
		logger.Panic("Error initializing JWT auth service", err)
	}

	return serviceJWT{
		usersService:    usersService,
		tokenSigningKey: []byte(config.TokenSigningKey),
		tokenExpiration: time.Duration(seconds) * time.Second,
	}
}

func (s serviceJWT) Login(user users.User) (domain.Auth, apierrors.APIError) {
	found, apiErr := s.usersService.Check(user)
	if apiErr != nil {
		logger.Error("Error logging user", apiErr)
		return domain.Auth{}, apiErr
	}

	expiresAt := time.Now().UTC().Add(s.tokenExpiration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		claimUserName:  found.UserName,
		claimExpiresAt: expiresAt,
	})

	signed, err := token.SignedString(s.tokenSigningKey)
	if err != nil {
		logger.Error("Error signing token", err)
		return domain.Auth{}, apierrors.NewInternalServerAPIError("error signing token")
	}

	return domain.Auth{
		Token:     signed,
		UserName:  found.UserName,
		ExpiresAt: expiresAt,
	}, nil
}

func (s serviceJWT) Validate(auth domain.Auth) (domain.Auth, apierrors.APIError) {
	parsed, err := jwt.Parse(auth.Token, func(value *jwt.Token) (interface{}, error) {
		if _, ok := value.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return s.tokenSigningKey, nil
	})
	if err != nil {
		logger.Error("Error parsing token", err)
		return domain.Auth{}, apierrors.NewInternalServerAPIError("error parsing token")
	}

	data := parsed.Claims.(jwt.MapClaims)
	exp, found := data[claimExpiresAt]
	if !found {
		err := errors.New("expires_at not found")
		logger.Error("Invalid token", err)
		return domain.Auth{}, apierrors.NewUnauthorizedAPIError("invalid token")
	}

	expiration, ok := exp.(float64)
	if !ok {
		err := errors.New("invalid expires_at")
		logger.Error("Invalid token", err)
		return domain.Auth{}, apierrors.NewUnauthorizedAPIError("invalid token")
	}

	expiresAt := int64(expiration)
	if time.Now().UTC().After(time.Unix(expiresAt, 0)) {
		err := errors.New("expired token")
		logger.Error("Invalid token", err)
		return domain.Auth{}, apierrors.NewUnauthorizedAPIError("expired token")
	}

	usr, found := data[claimUserName]
	if !found {
		err := errors.New("user_name not found")
		logger.Error("Invalid token", err)
		return domain.Auth{}, apierrors.NewUnauthorizedAPIError("invalid token")
	}

	userName, ok := usr.(string)
	if !ok {
		err := errors.New("invalid user_name")
		logger.Error("Invalid token", err)
		return domain.Auth{}, apierrors.NewUnauthorizedAPIError("invalid token")
	}

	if userName != auth.UserName {
		err := errors.New("invalid claim")
		logger.Error("Invalid token", err)
		return domain.Auth{}, apierrors.NewUnauthorizedAPIError("invalid token")
	}
	
	auth.ExpiresAt = expiresAt
	return auth, nil
}
