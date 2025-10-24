package main

import (
	"log"
	"payment-momo/internal/api"
	"payment-momo/internal/app"
	"payment-momo/internal/persistence"
	"payment-momo/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/viebiz/lit/env"
)

func main() {
	cfg, err := env.ReadAppConfig[app.Config]()
	if err != nil {
		log.Fatal(err)
	}
	orderRepo := persistence.NewOrderRepoInMemory()
	momoClient := persistence.NewMomoClient(cfg)
	paymentSvc := service.NewPaymentService(momoClient, orderRepo)

	r := gin.Default()
	api.RegisterRoutes(r, paymentSvc)

	log.Printf("🚀 MoMo API running at %s (base=%s)", cfg.AppPort, cfg.AppBaseURL)
	if err := r.Run(cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
