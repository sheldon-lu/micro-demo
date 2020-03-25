# micro-demo

beego-demo的micro-demo版本。
因为使用的gomicro准备使用微服务化，把用户、权限、后续其他都独立一个模块开发；
每个服务都是可以独立运行，go.mod的包并非micro-demo/开头的

~~~
``目录结构``
role-service RBAC权限模块服务，使用go-micro简单的增删改查实现demo
user-service 用户模块服务
mid-stage    类似中台服务，主要做网关层api，api鉴权等
roleMicro    用于测试rpc调用后台svc的role-service的客户端client，使用gin做api接口
RoleApi      用于测试rpc调用后台svc的role-service的客户端client
~~~
前后端分离，此为后端服务micro