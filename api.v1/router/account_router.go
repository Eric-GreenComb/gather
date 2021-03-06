package router

import (
	"github.com/banerwai/gather/api.v1/handler"
	"github.com/gin-gonic/gin"
)

// SetupV1AccountRoute account route
func SetupV1AccountRoute(g *gin.Engine) {
	r := g.Group("/api/v1")
	{
		r.GET("/account/:uid", handler.GetAccountBean)
		r.GET("/billing/:id", handler.GetBillingBean)
		r.GET("/billings/deal", handler.GetDealBillingBeans)
		r.GET("/billings/all", handler.GetBillingBeans)

		r.POST("account", handler.CreateAccountBean)
		r.POST("account/gen/:uid", handler.GenAccount)
		r.POST("billing", handler.CreateBillingBean)
		r.POST("billing/deal/:bid", handler.DealBilling)
		r.POST("billing/cancel/:bid", handler.CancelBilling)
	}
}
