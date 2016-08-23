package router

import (
	"github.com/banerwai/gather/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupV1AccountRoute(g *gin.Engine) {
	r := g.Group("/api/v1")
	{
		r.GET("/account/:uid", handler.GetAccountBean)
		r.GET("/billing/:id", handler.GetBillingBean)
		r.GET("/billings/deal", handler.GetDealBillingBeans)
		r.GET("/billings/all", handler.GetBillingBeans)

		r.POST("account/add", handler.CreateAccountBean)
		r.POST("account/gen/:uid", handler.GenAccount)
		r.POST("billing/add", handler.CreateBillingBean)
		r.POST("billing/deal/:bid", handler.DealBilling)
		r.POST("billing/cancel/:bid", handler.CancelBilling)
	}
}
