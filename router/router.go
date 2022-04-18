package router

import (
	"net/http"

	_ "apiserver/docs" // docs is generated by Swag CLI, you have to import it.
	"apiserver/handler/sd"
	"apiserver/handler/user"
	"apiserver/router/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 加载路由、中间件、 handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 中间件
	/*gin.Recovery()：在处理某些请求时可能因为程序bug或者其他异常情况导致程序panic,这时候为了
	不影响下一次请求的调用，需要通过gin.Recovery()来恢复API服务器*/
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache) //middleware.NoCache:
	g.Use(middleware.Options) //middleware.Options:
	g.Use(middleware.Secure)  //middleware.Secure:
	g.Use(mw...)              //
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	// api for authentication functionalities
	g.POST("/login", user.Login)

	// 用户路由设置
	/*在RESTful API开发中，API经常会变动，为了兼容老的API,引入了版本的概念，比如上例中的v1/user,说明该API版本是v1。
	很多RESTful API最佳实践文章中均建议使用版本控制，笔者这里也建议对API使用版本控制。*/
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)       //创建用户
		u.DELETE("/:id", user.Delete) //删除用户
		u.PUT("/:id", user.Update)    //更新用户
		u.GET("", user.List)          //用户列表
		u.GET("/:username", user.Get) //获取指定用户的信息
	}
	//Sd分组主要用来检查API Server的状态：健康状况、服务器硬盘、CPU和内存使用量。
	//健康检查handlers
	//设置路由组
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
