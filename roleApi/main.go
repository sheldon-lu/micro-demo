package main

import (
    `fmt`
    `log`

    "github.com/gin-gonic/gin"
    "micro-demo/roleApi/handler"

    `net/http`
    `time`
)

func main() {
    // create new web service
    //service := web.NewService(
    //    web.Name("role-service.v1"),
    //    web.Version("latest"),
    ////    web.Address(":8089"),
    //)
    //
    //// initialise service
    //if err := service.Init(
    //    web.Action(
    //        func(c *cli.Context) {
    //            // 初始化handler
    //            handler.Init()
    //        }),
    //); err != nil {
    //    log.Fatal(err)
    //}


    handler.Init()

    // register html handler
    // service.Handle("/", http.FileServer(http.Dir("html")))

    // register call handler
    // Create RESTful handler (using Gin)
    // Init new(User)
    role := new(handler.Roles)
    router := gin.Default()

    v1 := router.Group("/v1")
    {
        v1.GET("/role", func(c *gin.Context){fmt.Println("33")})
        v1.POST("/role/create", role.Create)
    }

    server := &http.Server{
        Addr:           ":8080",
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    log.Printf("begin listen Address %v", ":8080")
    server.ListenAndServe()


    //service.Handle("/", router)

    // run service
    //if err := service.Run(); err != nil {
    //    log.Fatal(err)
    //}
}

