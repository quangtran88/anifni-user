package ports

import "github.com/quangtran88/anifni-user/core/domain"

type UserService interface {
	Get(id domain.ID) (domain.User, error)
	Register(input domain.RegisterUserInput) (domain.User, error)
}
