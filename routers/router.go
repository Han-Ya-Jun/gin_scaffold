package routers

import (
	_ "github.com/Han-Ya-Jun/gin_scaffold/docs"
	"github.com/Han-Ya-Jun/gin_scaffold/middleware/access"
	"github.com/Han-Ya-Jun/gin_scaffold/middleware/logging"
	"github.com/Han-Ya-Jun/gin_scaffold/pkg/def"
	v1 "github.com/Han-Ya-Jun/gin_scaffold/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

/*
* @Author:hanyajun
* @Date:2019/6/06 16:50
* @Name:routers
* @Function:
 */

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(logging.RequestLog())
	r.Use(gin.Recovery())
	r.Use(access.Control())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//健康检查
	r.GET("/health", HealthHandler)
	ApiV1 := r.Group("gin_scaffold/api/v1/mongo")

	{
		//add user
		ApiV1.POST("/user", v1.AddUser)
		////select user
		//ApiV1.GET("/user/:uid", v1.OrganizationGet)
		////update user
		//ApiV1.PUT("/user/:uid", v1.WaybillPush)
		////delete user
		//ApiV1.POST("/user/:uid", v1.WaybillBatchGet)
		////find users with pages
		//ApiV1.GET("/users", v1.HubsRegister)
	}
	return r
}

type Health struct {
	GoVersion string
}

func HealthHandler(c *gin.Context) {
	h := Health{
		GoVersion: def.GO_VERSION,
	}
	c.JSON(
		200,
		h,
	)

}
