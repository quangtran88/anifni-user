package ports

import (
	"context"
	"github.com/quangtran88/anifni-user/core/domain"
)

type UserService interface {
	Get(ctx context.Context, id domain.ID) (domain.User, error)
	CheckDuplicated(ctx context.Context, email string) (bool, error)
	Create(ctx context.Context, in domain.CreateUserInput) (domain.User, error)
}
