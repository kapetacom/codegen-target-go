module github.com/kapeta/todo

go 1.21.7

toolchain go1.21.8

require (
	github.com/kapetacom/sdk-go-auth-jwt v0.0.2
	github.com/kapetacom/sdk-go-config v0.1.3
	github.com/kapetacom/sdk-go-nosql-mongodb v0.0.5
	github.com/kapetacom/sdk-go-rabbitmq v0.1.0
	github.com/kapetacom/sdk-go-rest-client v0.0.2
	github.com/kapetacom/sdk-go-rest-server v0.1.3
	github.com/labstack/echo/v4 v4.11.4
)

require (
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang-jwt/jwt/v5 v5.2.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.5 // indirect
	github.com/kapetacom/schemas/packages/go v0.0.0-20240209083259-f5ce079d8abc // indirect
	github.com/klauspost/compress v1.17.6 // indirect
	github.com/labstack/echo-jwt/v4 v4.2.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/lestrrat-go/backoff/v2 v2.0.8 // indirect
	github.com/lestrrat-go/blackmagic v1.0.2 // indirect
	github.com/lestrrat-go/httpcc v1.0.1 // indirect
	github.com/lestrrat-go/iter v1.0.2 // indirect
	github.com/lestrrat-go/jwx v1.2.28 // indirect
	github.com/lestrrat-go/option v1.0.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rabbitmq/amqp091-go v1.9.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/wagslane/go-rabbitmq v0.12.4 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.mongodb.org/mongo-driver v1.13.1 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Pending PR against upstream: https://github.com/wagslane/go-rabbitmq/pull/152
replace github.com/wagslane/go-rabbitmq => github.com/kapetacom/go-rabbitmq v1.0.0
