package service

//go:generate mockery --name RateLimiterServiceInterface --outpkg mock --output mock --filename ratelimite.go --with-expecter=true

type RateLimiterServiceInterface interface {
	AllowRequest(ip, token string) bool
}

type RateLimiterService struct {
	RateLimiterRepo RateLimiterServiceInterface
}

func NewRateLimiterService(repo RateLimiterServiceInterface) *RateLimiterService {
	return &RateLimiterService{
		RateLimiterRepo: repo,
	}
}

func (rls *RateLimiterService) AllowRequest(ip, token string) bool {
	return rls.RateLimiterRepo.AllowRequest(ip, token)
}
