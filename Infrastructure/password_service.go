package Infrastructure

import "golang.org/x/crypto/bcrypt"

type PasswordServiceBcrypt struct{}

func (p *PasswordServiceBcrypt) GenerateHash(plain string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)

	return string(hashedPass), err
}

func (p *PasswordServiceBcrypt) ComparePassword(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	return err == nil
}
