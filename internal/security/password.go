package security

import "golang.org/x/crypto/bcrypt"

// Hash generates a bcrypt hash for a password string
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compares the hashed password and the clear text password
func VerifyPassword(hashedPassword, clearPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(clearPassword))
}
