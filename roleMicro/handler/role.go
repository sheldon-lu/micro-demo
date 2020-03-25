package handler

import (
    `context`
    `fmt`
    rolepb "micro-demo/role-service/protos/roleService"
    `github.com/gin-gonic/gin`
    `github.com/go-acme/lego/v3/log`
    `github.com/micro/go-micro/client`
)

type Roles struct {
}

var (
    cli1 rolepb.RoleService
)

type Result struct {
    Result     bool   `json:"result"`
    Message    string `json:"message"`
    ResultCode int    `json:"result_code,omitempty"`
}

// setup Greeter Server Client Init
func Init() {
    cli1 = rolepb.NewRoleService("role-service", client.DefaultClient)
}

func (us *Roles) Create(c *gin.Context) {
    name := "1234567"
    pid := "1"
    status := int64(1)
    remark := "2"

    resp, err := cli1.CreateRole(context.TODO(), &rolepb.CreateRoleRequest{
        RoleName:   name,
        RoleRemark: remark,
        RoleStatus: status,
        RolePid:    pid,
    })
    if err != nil {
        Result := Result{
            Result:     false,
            Message:    fmt.Sprintln(err),
            ResultCode: 40000,
        }
        log.Println(err)
        c.JSON(500, Result)
        return
    }

    c.JSON(200, resp)
}

func (us *Roles) Update(c *gin.Context) {
    name := "1235"
    pid := "1"
    status := int64(1)
    remark := "333"

    resp, err := cli1.UpdateRole(context.TODO(), &rolepb.UpdateRoleRequest{
        RoleName:   name,
        RoleRemark: remark,
        RoleStatus: status,
        RolePid:    pid,
    })

    if err != nil {
        c.JSON(500, c.Error(err))
        return
    }

    c.JSON(200, resp)
}

func (us *Roles) GetRoleList(c *gin.Context) {

    resp, err := cli1.GetRoles(context.TODO(), &rolepb.GetRoleRequest{})

    if err != nil {
        c.JSON(500, c.Error(err))
        return
    }

    c.JSON(200, resp)
}

func (us *Roles) DeleteDemo(c *gin.Context) {
    id := "2"

    resp, err := cli1.DeleteRole(context.TODO(), &rolepb.DeleteRoleRequest{
        RoleId: id,
    })

    if err != nil {
        c.JSON(500, c.Error(err))
        return
    }

    c.JSON(200, resp)
}

func (us *Roles) RoleAddAccess(c *gin.Context) {
    role_id := int64(2)
    pas := "1_1_0,3_2_1,4_3_3,8_3_2,13_3_2,14_3_2"

    resp, err := cli1.RoleAddAccess(context.TODO(), &rolepb.RoleAddAccessRequest{
        RoleId: role_id,
        Params: pas,
    })

    if err != nil {
        c.JSON(500, c.Error(err))
        return
    }

    c.JSON(200, resp)
}
