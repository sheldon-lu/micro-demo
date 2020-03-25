package main

import (
    `fmt`
    "github.com/micro/go-micro/util/log"
    "github.com/micro/go-micro/web"
    "github.com/gin-gonic/gin"
    "micro-demo/roleMicro/handler"
    "github.com/micro/cli"
)

func main() {
    // create new web service
    service := web.NewService(
        web.Name("role-service.v1"),
        web.Version("latest"),
        web.Address(":8089"),
    )

    // initialise service
    if err := service.Init(
        web.Action(
            func(c *cli.Context) {
                // 初始化handler
                handler.Init()
            }),
    ); err != nil {
        log.Fatal(err)
    }

    // register html handler
    // service.Handle("/", http.FileServer(http.Dir("html")))

    // register call handler
    // Create RESTful handler (using Gin)
    // Init new(User)
    role := new(handler.Roles)
    router := gin.Default()

    v1 := router.Group("/v1")
    {
        v1.GET("/role", func(c *gin.Context) { fmt.Println("33") })
        v1.POST("/role/create", role.Create)
        v1.POST("/role/update", role.Update)
        v1.POST("/role/list", role.GetRoleList)
        v1.POST("/role/delete", role.DeleteDemo)
        v1.POST("/role/rladdas", role.RoleAddAccess)
    }

    service.Handle("/", router)

    // run service
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
