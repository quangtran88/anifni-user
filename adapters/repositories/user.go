package repositories

import (
	"context"
	"github.com/quangtran88/anifni-user/core/domain"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo UserRepository) FindById(ctx context.Context, id domain.ID) (*domain.User, error) {
	return &domain.User{}, nil
}

func (repo UserRepository) FindByPId(ctx context.Context, pid domain.PID) (*domain.User, error) {
	return &domain.User{}, nil
}

func (repo UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (repo UserRepository) Create(ctx context.Context, user domain.User) (domain.ID, error) {
	return "", nil
}
