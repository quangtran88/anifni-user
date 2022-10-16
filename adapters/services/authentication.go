package serviceAdapter

type AuthenticationService struct {
}

func NewAuthenticationService() *AuthenticationService {
	return &AuthenticationService{}
}

func (srv AuthenticationService) CheckEmailOTP(code string, email string) (bool, error) {
	return true, nil
}
