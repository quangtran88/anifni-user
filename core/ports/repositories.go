package ports

import "github.com/quangtran88/anifni-user/core/domain"

type UserRepository interface {
	FindById(id domain.ID) (domain.User, error)
}
