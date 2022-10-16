package serviceAdapter

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (srv TokenService) GetToken(domain string, length int) (string, error) {
	return "", nil
}
