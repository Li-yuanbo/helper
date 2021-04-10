package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	. "helper/handler"
)

func handler(r *gin.Engine){
	r.POST("/user", RegisterUser) //注册用户
}

func main() {
	//默认路由
	r := gin.Default()
	store := cookie.NewStore([]byte("user"))
	store.Options(sessions.Options{
		MaxAge: 24 * 60 * 60,
		Path:   "/",
	})
	//在路由上加入session中间件
	r.Use(sessions.Sessions("mysession", store))
	//加载handler
	handler(r)
	//启动项目
	r.Run(":8080")
}
