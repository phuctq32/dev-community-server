package hasher

import "golang.org/x/crypto/bcrypt"

type bcryptHash struct {
	cost int
}

func NewBcryptHash(cost int) *bcryptHash {
	return &bcryptHash{cost: cost}
}

func (h *bcryptHash) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (h *bcryptHash) ComparePassword(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
