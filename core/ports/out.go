package ports

import "github.com/quangtran88/anifni-user/core/domain"

type UserRepository interface {
	FindById(id domain.ID) (domain.User, error)
	Create(user domain.User) (domain.ID, error)
}

type AuthenticationService interface {
	CheckEmailOTP(code string, email string) (bool, error)
}

type TokenService interface {
	GetToken(domain string, length int) (string, error)
}
