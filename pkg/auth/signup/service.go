package signup

import "golang.org/x/crypto/bcrypt"

type Storage interface {
	NewUser(User) error
}

type service struct {
	s Storage
}

func NewSignUpService(s Storage) *service {
	return &service{s}
}

func (service *service) SignUp(user User) error {
	hashed, error := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if error != nil {
		return error
	}

	user.Password = string(hashed)

	if error := service.s.NewUser(user); error != nil {
		return error
	}

	return nil
}
