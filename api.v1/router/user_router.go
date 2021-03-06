package router

import (
	"github.com/banerwai/gather/api.v1/handler"
	"github.com/gin-gonic/gin"
)

// SetupV1UserRoute user route
func SetupV1UserRoute(g *gin.Engine) {
	r := g.Group("/api/v1")
	{
		r.GET("/user/:email", handler.GetUserByEmail)
		r.GET("/userid/:id", handler.GetUserByID)

		r.POST("user/add", handler.CreateBeanUser)
		r.POST("user/reset", handler.ResetPwd)
		r.POST("user/active", handler.ActiveUser)

	}
}
