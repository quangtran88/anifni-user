package ports

import (
	"context"
	"github.com/quangtran88/anifni-user/core/domain"
)

type UserRepository interface {
	FindById(ctx context.Context, id domain.ID) (*domain.User, error)
	FindByPId(ctx context.Context, pid domain.PID) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.ID, error)
}

type AuthenticationService interface {
	CheckEmailOTP(ctx context.Context, code string, email string) (bool, error)
}

type TokenService interface {
	GetToken(ctx context.Context, domain string, length int) (string, error)
}

type HashGenerator interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
