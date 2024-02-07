package auth

import (
	"github.com/kapetacom/sdk-go-auth-jwt/middleware"
	"github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

func AddJWTMiddleware(config providers.ConfigProvider) []echo.MiddlewareFunc {
	return middleware.JWTMiddlewareFromConfig("auth", config)
}
