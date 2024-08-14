package interfaces

import "github.com/dgrijalva/jwt-go"

type IJWTService interface {
	GenerateToken(string) (string, error)
	ValidateToken(string) (*jwt.Token, error)
	GetUserName(*jwt.Token) (string, error)
}
