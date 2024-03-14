module github.com/kapeta/users

go 1.21.6

require (
	github.com/kapetacom/sdk-go-config v0.1.3
	github.com/kapetacom/sdk-go-rabbitmq v0.1.0
	github.com/kapetacom/sdk-go-rest-server v0.2.2
	github.com/labstack/echo/v4 v4.11.4
)

require (
	github.com/ggicci/httpin v0.16.0 // indirect
	github.com/ggicci/owl v0.7.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.5 // indirect
	github.com/kapetacom/schemas/packages/go v0.0.0-20240209083259-f5ce079d8abc // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/rabbitmq/amqp091-go v1.9.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/wagslane/go-rabbitmq v0.12.4 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Pending PR against upstream: https://github.com/wagslane/go-rabbitmq/pull/152
replace github.com/wagslane/go-rabbitmq => github.com/kapetacom/go-rabbitmq v1.0.0
