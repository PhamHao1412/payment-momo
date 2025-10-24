package api

import (
	"embed"
	"net/http"
	"payment-momo/internal/service"

	"github.com/gin-gonic/gin"
)

//go:embed static/*.html
var staticFS embed.FS

func RegisterRoutes(r *gin.Engine, paymentSvc *service.PaymentService) {
	h := NewHandler(paymentSvc)

	// Serve embedded index.html
	r.GET("/", func(c *gin.Context) {
		b, _ := staticFS.ReadFile("static/index.html")
		c.Data(http.StatusOK, "text/html; charset=utf-8", b)
	})

	r.GET("/momo/return", func(c *gin.Context) {
		b, _ := staticFS.ReadFile("static/return.html")
		c.Data(http.StatusOK, "text/html; charset=utf-8", b)
	})

	v1 := r.Group("/api/v1/payment/momo")
	{
		v1.POST("/create", h.CreatePayment)
		v1.POST("/ipn", h.HandleIPN)
		v1.GET("/check-status", h.CheckStatus)
		v1.POST("/update-order", h.UpdateOrder)
	}
}
