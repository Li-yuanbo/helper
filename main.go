package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	. "helper/handler"
	"log"
)

func handler(r *gin.Engine) {
	//user_info
	r.POST("/user", RegisterUser)        //注册用户
	r.GET("/users", GetUserInfos)        //分页获取用户基础信息
	r.GET("/user/login", LoginUser)      //用户登录
	r.DELETE("/user/login", UnLoginUser) //退出登录

	//organization
	r.POST("/org", RegisterOrganization) //注册志愿者机构
	r.GET("/org/login", LoginOrg)        //志愿者机构登录
	r.DELETE("/org/login", UnLoginOrg)   //志愿者机构退出登录

	//activity
	r.POST("/activity", PublishActivity) //发布活动
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
	err := r.Run(":8080")
	if err != nil {
		log.Println("run gin err: ", err)
	}
}
