package usecases

import (
	"github.com/Abzaek/clean-arch/domain"
	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateToken(user domain.User) (string, error)
	ValidateToken(tokenStr string) (*jwt.MapClaims, error)
}
