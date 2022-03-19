package signup

type Service interface {
	SignUp(User) error
}

type Storage interface {
	SignUp(User) error
}

type service struct {
	s Storage
}

func NewSignUpService(s Storage) Service {
	return &service{s}
}

func (service *service) SignUp(user User) error {
	if error := service.s.SignUp(user); error != nil {
		return error
	}

	return nil
}
