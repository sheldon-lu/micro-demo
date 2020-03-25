package handler

import (
    `context`
    rolepb "micro-demo/role-service/protos/roleService"
    `github.com/gin-gonic/gin`
    `github.com/micro/go-micro`
)

type Roles struct {
}

var (
    cli1 rolepb.RoleService
)


// setup Greeter Server Client Init
func Init() {

    //创建服务
    service := micro.NewService(
        micro.Name("rolesssService"),
        micro.Version("latest"),
    )
    service.Init()
    cli1 = rolepb.NewRoleService("role-service", service.Client())
}

func (us *Roles) Create(c *gin.Context) {
    name := "123"
    pid := "1"
    status := int64(1)
    remark := "324"

    resp, err := cli1.CreateRole(context.TODO(), &rolepb.CreateRoleRequest{
        RoleName: name,
        RoleRemark: remark,
        RoleStatus: status,
        RolePid: pid,
    })
    if err != nil {
        c.JSON(500, c.Error(err))
        return
    }

    c.JSON(200, resp)
}

