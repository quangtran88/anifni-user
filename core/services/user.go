package services

import (
	"context"
	"github.com/quangtran88/anifni-user/core/domain"
	"github.com/quangtran88/anifni-user/core/ports"
)

type UserService struct {
	repo ports.UserRepository
	hash ports.HashGenerator
}

func NewUserService(
	repo ports.UserRepository,
	hash ports.HashGenerator,
) *UserService {
	return &UserService{repo, hash}
}

func (srv UserService) Get(ctx context.Context, id domain.ID) (domain.User, error) {
	user, err := srv.repo.FindById(ctx, id)
	return *user, err
}

func (srv UserService) CheckDuplicated(ctx context.Context, email string) (bool, error) {
	existed, err := srv.repo.FindByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	if existed == nil {
		return true, nil
	}
	return false, nil
}

func (srv UserService) Create(ctx context.Context, in domain.CreateUserInput) (domain.User, error) {
	hashedPassword, err := srv.hash.HashPassword(in.Password)
	if err != nil {
		return domain.User{}, err
	}

	in.Password = hashedPassword
	user := domain.NewUser(in)

	createdId, err := srv.repo.Create(ctx, *user)
	if err != nil {
		return domain.User{}, err
	}

	user.Id = createdId
	return *user, nil
}
