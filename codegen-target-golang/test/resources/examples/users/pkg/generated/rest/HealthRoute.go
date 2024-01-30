// GENERATED SOURCE - DO NOT EDIT
package rest

import (
	sdkgoconfig "github.com/kapetacom/golang-language-target/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

// CreateHealth creates health endpoint for the service
func CreateHealth(e *echo.Echo, cfg sdkgoconfig.ConfigProvider) {
	e.GET("/.kapeta/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
}
