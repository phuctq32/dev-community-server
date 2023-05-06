package hasher

type MyHash interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) bool
}
