package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"school-walker/walker/middlerware"
	"school-walker/walker/pkg/setting"
	"school-walker/walker/routers/v1/leave"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlerware.ErrorPrinter())
	app, err := setting.Cfg.GetSection("app")
	if err != nil {
		log.Fatal("load app config error:", err.Error())
	}

	gin.SetMode(app.Key("RUN_MODE").String())
	// gin.SetMode(gin.ReleaseMode)

	api := r.Group("/api")
	{
		leaveV1 := api.Group("/v1/leave")
		{
			// check user leave authority
			leaveV1.GET("/verification", leave.AuthCheck)
			// save user leave info
			leaveV1.POST("/save", leave.SaveAuth)
		}
	}

	// view := r.Group("/view")
	// {
	// 	viewV1 := view.Group("/v1/leave")
	// 	viewV1.GET("/first", nil)
	// 	viewV1.GET("/page", nil)
	// }
	return r
}
