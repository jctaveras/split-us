package signup

type Storage interface {
	SignUp(User) error
}

type service struct {
	s Storage
}

func NewSignUpService(s Storage) *service {
	return &service{s}
}

func (service *service) SignUp(user User) error {
	if error := service.s.SignUp(user); error != nil {
		return error
	}

	return nil
}
