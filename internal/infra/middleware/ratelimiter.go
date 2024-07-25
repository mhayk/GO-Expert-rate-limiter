package middleware

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/mhayk/GO-Expert-rate-limiter/internal/application"
)

type RateLimiterMiddleware struct {
	RateLimiterApp application.RateLimiterServiceInterface
}

func NewRateLimiterMiddleware(rateLimiterApp application.RateLimiterServiceInterface) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		RateLimiterApp: rateLimiterApp,
	}
}

func (rlm *RateLimiterMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := GetIP(r)
		token := r.Header.Get("API_KEY")

		log.Printf("IP: %s, Token: %s", ip, token)

		if rlm.RateLimiterApp.AllowRequest(ip, token) {
			next.ServeHTTP(w, r)
		} else {
			log.Printf("Request denied for IP: %s, Token: %s", ip, token)
			http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
		}
	})
}

func GetIP(r *http.Request) string {
	ip := r.RemoteAddr
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		ip = strings.Split(forwarded, ",")[0]
	} else {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			return ""
		}
		ip = host
	}
	return ip
}
