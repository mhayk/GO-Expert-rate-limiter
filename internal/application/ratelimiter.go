package application

import (
	"github.com/mhayk/GO-Expert-rate-limiter/internal/domain/service"
)

//go:generate mockery --name RateLimiterServiceInterface --outpkg mock --output mock --filename ratelimite.go --with-expecter=true

type RateLimiterServiceInterface interface {
	AllowRequest(ip, token string) bool
}

type RateLimiterApp struct {
	RateLimiterService *service.RateLimiterService
}

func NewRateLimiterApp(rateLimiterService *service.RateLimiterService) *RateLimiterApp {
	return &RateLimiterApp{
		RateLimiterService: rateLimiterService,
	}
}

func (app *RateLimiterApp) AllowRequest(ip, token string) bool {
	return app.RateLimiterService.AllowRequest(ip, token)
}
