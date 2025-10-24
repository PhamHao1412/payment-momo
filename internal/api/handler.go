package api

import (
	"net/http"

	"payment-momo/internal/model"
	"payment-momo/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	paymentSvc *service.PaymentService
}

func NewHandler(svc *service.PaymentService) *Handler {
	return &Handler{paymentSvc: svc}
}

func (h *Handler) CreatePayment(c *gin.Context) {
	var req model.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	res, err := h.paymentSvc.CreatePayment(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) HandleIPN(c *gin.Context) {
	var cb model.MomoCallback
	if err := c.ShouldBindJSON(&cb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := h.paymentSvc.HandleIPN(c, cb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"resultCode": 0, "message": "IPN received"})
}

func (h *Handler) CheckStatus(c *gin.Context) {
	orderID := c.Query("orderId")
	requestID := c.Query("requestId")
	if orderID == "" || requestID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing orderId/requestId"})
		return
	}

	res, err := h.paymentSvc.CheckStatus(c, orderID, requestID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	var req struct {
		OrderId string `json:"orderId"`
		Status  string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	if err := h.paymentSvc.UpdateOrder(req.OrderId, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order updated"})
}
