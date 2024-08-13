package usecases

type IPassword interface {
	HashPassword(string) (string, error)
}
