package router

import (
	"github.com/banerwai/gather/api.v1/handler"
	"github.com/gin-gonic/gin"
)

func SetupV1ContactRoute(g *gin.Engine) {
	r := g.Group("/api/v1")
	{
		r.GET("/contact/tpl/:tplname", handler.GetContactTpl)
		r.GET("/contact/get/:id", handler.GetContactBean)
		r.GET("/contact/get/:id/status", handler.GetContactSignStatusEnum)
		r.GET("/contact/client/:email", handler.GetClientContactBeans)
		r.GET("/contact/freelancer/:email", handler.GetFreelancerContactBeans)

		r.POST("contact/add", handler.CreateContactBean)
		r.POST("contact/update", handler.UpdateContact)
		r.POST("contact/sign/client", handler.ClientSignContact)
		r.POST("contact/sign/freelance", handler.FreelancerSignContact)
		r.POST("contact/deal", handler.DealContact)
	}
}
