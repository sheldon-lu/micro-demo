module micro-demo/roleApi

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	micro-demo/role-service v0.0.0-00010101000000-000000000000
)

replace micro-demo/role-service => ../role-service
