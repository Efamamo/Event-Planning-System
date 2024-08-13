package usecases

type IJWTService interface {
	GenerateToken(string) (string, error)
	GetUserName(t string) (string, error)
}
