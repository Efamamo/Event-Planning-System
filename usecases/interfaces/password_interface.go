package interfaces

type IPassword interface {
	HashPassword(string) (string, error)
	ComparePassword(euPassword string, uPassword string) (bool, error)
}
