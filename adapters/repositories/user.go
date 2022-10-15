package repositories

import "github.com/quangtran88/anifni-user/core/domain"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo UserRepository) FindById(id domain.ID) (domain.User, error) {
	return domain.User{Id: "U001", Name: "Randy Tran"}, nil
}
