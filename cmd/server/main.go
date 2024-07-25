package main

import (
	"log"
	"net/http"

	"github.com/mhayk/GO-Expert-rate-limiter/internal/application"
	"github.com/mhayk/GO-Expert-rate-limiter/internal/domain/service"
	"github.com/mhayk/GO-Expert-rate-limiter/internal/infra/config"
	"github.com/mhayk/GO-Expert-rate-limiter/internal/infra/limiter"
	"github.com/mhayk/GO-Expert-rate-limiter/internal/infra/middleware"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	log.Printf("Redis configuration: host=%s, port=%s, password=%s", cfg.RedisHost, cfg.RedisPort, cfg.RedisPassword)

	redisAddr := cfg.RedisHost + ":" + cfg.RedisPort

	ipStore := limiter.NewRedisStore(redisAddr, cfg.RedisPassword)
	tokenStore := limiter.NewRedisStore(redisAddr, cfg.RedisPassword)

	rateLimiter := limiter.NewLimiter(
		cfg.IPMaxRequestsPerSecond,
		cfg.IPBlockDurationSeconds,
		cfg.TokenMaxRequestsPerSecond,
		cfg.TokenBlockDurationSeconds,
		ipStore,
		tokenStore,
	)

	rateLimiterService := service.NewRateLimiterService(rateLimiter)

	rateLimiterApp := application.NewRateLimiterApp(rateLimiterService)

	rateLimiterMiddleware := middleware.NewRateLimiterMiddleware(rateLimiterApp)

	mux := http.NewServeMux()
	mux.Handle("/", rateLimiterMiddleware.Handler(http.HandlerFunc(handler)))

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
