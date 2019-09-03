package routes

import (
	"accbase/app/Controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter()  *gin.Engine{
	r := gin.Default()

	r.Static("/public", "./public") // 静态文件服务
	r.LoadHTMLGlob("views/**/*") // 载入html模板目录
	// web前端路由
	r.GET("/", Controllers.Home)
	r.GET("/about", Controllers.About)
	r.GET("/post", Controllers.Post)
	r.GET("/user", Controllers.User)

	// web 后端路由
	r.GET("/admin", Controllers.Admin)
	r.POST("/api/token", Controllers.GetToken)

	// api 路由
	InitApi(r)

	return r
}