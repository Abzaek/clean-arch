package usecases

type PasswordService interface {
	ComparePassword(hashed string, plain string) bool
	GenerateHash(plain string) (string, error)
}
