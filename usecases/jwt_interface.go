package usecases

import "github.com/dgrijalva/jwt-go"

type IJWTService interface {
	GenerateToken(username string) (string, error)
	ValidateToken(t string) (*jwt.Token, error)
}
