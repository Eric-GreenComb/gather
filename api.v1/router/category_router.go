package router

import (
	"github.com/banerwai/gather/api.v1/handler"
	"github.com/gin-gonic/gin"
)

// SetupV1CategoryRoute category route
func SetupV1CategoryRoute(g *gin.Engine) {
	r := g.Group("/api/v1")
	{
		r.GET("/loadcategory/:path", handler.LoadCategory)
		r.GET("/category", handler.GetCategoriesBean)
		r.GET("/category/:sn", handler.GetSubCategoriesBean)
	}
}
