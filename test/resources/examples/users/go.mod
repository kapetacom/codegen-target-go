module github.com/kapeta/todo

go 1.21.6

require (
    github.com/kapetacom/sdk-go-config v0.1.1
    github.com/labstack/echo/v4 v4.11.3
    github.com/kapetacom/sdk-go-auth-jwt v0.0.2
    github.com/kapetacom/sdk-go-rest-server v0.1.0
    github.com/kapetacom/sdk-go-rabbitmq v0.1.0
)

// Pending PR against upstream: https://github.com/wagslane/go-rabbitmq/pull/152
replace github.com/wagslane/go-rabbitmq => github.com/kapetacom/go-rabbitmq v1.0.0
