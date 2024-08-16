package Infrastructure

import (
	"errors"
	"time"

	usecases "github.com/Abzaek/clean-arch/Usecases"
	"github.com/Abzaek/clean-arch/domain"
	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	JwtKey  []byte
	Service usecases.UserService
}

func NewJwtService(jwtKey string) *JwtService {
	return &JwtService{
		JwtKey: []byte(jwtKey),
	}
}

func (s *JwtService) ValidateToken(t string) (*jwt.MapClaims, error) {

	token, err := jwt.ParseWithClaims(t, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid method")
		}

		return s.JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if exp, ok := (*claims)["exp"].(float64); ok {
		expirationTime := time.Unix(int64(exp), 0)

		if time.Now().After(expirationTime) {
			return nil, errors.New("token expired")
		}
	}

	var existingUser *domain.User

	existingUser, err = (s.Service).Find((*claims)["user_id"].(string))

	if err != nil {
		return &jwt.MapClaims{}, err
	}

	(*claims)["user_id"] = existingUser.Role

	return claims, nil
}

func (s *JwtService) GenerateToken(user domain.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(25 * time.Hour),
	})

	signedToken, err := token.SignedString(s.JwtKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
