package services

import (
	"errors"
	"fmt"
	"github.com/quangtran88/anifni-user/core/domain"
	"github.com/quangtran88/anifni-user/core/ports"
	"github.com/quangtran88/anifni-user/utils"
)

type UserService struct {
	repo     ports.UserRepository
	authSrv  ports.AuthenticationService
	tokenSrv ports.TokenService
}

func NewUserService(
	repo ports.UserRepository,
	authSrv ports.AuthenticationService,
	tokenSrv ports.TokenService,
) *UserService {
	return &UserService{repo: repo, authSrv: authSrv, tokenSrv: tokenSrv}
}

func (srv UserService) Get(id domain.ID) (domain.User, error) {
	return srv.repo.FindById(id)
}

func (srv UserService) Register(input domain.RegisterUserInput) (domain.User, error) {
	err := srv.checkOTP(input.OTPCode, input.Email)
	if err != nil {
		return domain.User{}, err
	}

	pid, err := srv.tokenSrv.GetToken(domain.UserTokenDomain, domain.UserTokenLength)
	if err != nil {
		return domain.User{}, err
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		PId:       domain.PID(pid),
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Name:      fmt.Sprintf("%s %s", input.FirstName, input.LastName),
		Password:  hashedPassword,
	}
	id, err := srv.repo.Create(user)
	if err != nil {
		return domain.User{}, err
	}

	user.Id = id
	return user, nil
}

func (srv UserService) checkOTP(code string, email string) error {
	ok, err := srv.authSrv.CheckEmailOTP(code, email)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("OTP code invalid")
	}
	return nil
}
