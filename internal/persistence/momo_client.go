package persistence

import (
	"fmt"
	"payment-momo/internal/app"
	"strconv"
	"strings"
	"time"

	"payment-momo/internal/model"
	"payment-momo/pkg"
)

type MomoClient struct{ cfg app.Config }

func NewMomoClient(cfg app.Config) *MomoClient { return &MomoClient{cfg: cfg} }

func (m *MomoClient) Config() app.Config { return m.cfg }

func (m *MomoClient) CreatePayment(req model.MomoCreateRequest) (*model.MomoCreateResponse, error) {
	raw := buildCreateRawSignature(req)
	req.Signature = pkg.HmacSHA256(raw, m.cfg.SecretKey)
	return pkg.HttpPostJSON[model.MomoCreateResponse]("https://test-payment.momo.vn/v2/gateway/api/create", req)
}

func (m *MomoClient) VerifyCallbackSignature(cb model.MomoCallback) bool {
	raw := buildCallbackRawSignature(cb, m.cfg.AccessKey)
	expected := pkg.HmacSHA256(raw, m.cfg.SecretKey)
	return strings.EqualFold(expected, cb.Signature)
}

func (m *MomoClient) QueryStatus(orderID, requestID string) (*model.QueryStatusResponse, error) {
	raw := fmt.Sprintf("accessKey=%s&orderId=%s&partnerCode=%s&requestId=%s", m.cfg.AccessKey, orderID, m.cfg.PartnerCode, requestID)
	req := model.QueryStatusRequest{
		PartnerCode: m.cfg.PartnerCode,
		AccessKey:   m.cfg.AccessKey,
		RequestId:   requestID,
		OrderId:     orderID,
		Signature:   pkg.HmacSHA256(raw, m.cfg.SecretKey),
	}
	return pkg.HttpPostJSON[model.QueryStatusResponse]("https://test-payment.momo.vn/v2/gateway/api/query", req)
}

// signature builders (order matters per MoMo docs)
func buildCreateRawSignature(p model.MomoCreateRequest) string {
	pairs := []string{
		"accessKey=" + p.AccessKey,
		"amount=" + p.Amount,
		"extraData=" + p.ExtraData,
		"ipnUrl=" + p.IpnUrl,
		"orderId=" + p.OrderId,
		"orderInfo=" + p.OrderInfo,
		"partnerCode=" + p.PartnerCode,
		"redirectUrl=" + p.RedirectUrl,
		"requestId=" + p.RequestId,
		"requestType=" + p.RequestType,
	}
	return strings.Join(pairs, "&")
}

func buildCallbackRawSignature(cb model.MomoCallback, accessKey string) string {
	pairs := []string{
		"accessKey=" + accessKey,
		"amount=" + cb.Amount,
		"extraData=" + cb.ExtraData,
		"message=" + cb.Message,
		"orderId=" + cb.OrderId,
		"orderInfo=" + cb.OrderInfo,
		"orderType=" + cb.OrderType,
		"partnerCode=" + cb.PartnerCode,
		"payType=" + cb.PayType,
		"requestId=" + cb.RequestId,
		"responseTime=" + cb.ResponseTime,
		"resultCode=" + cb.ResultCode,
		"transId=" + cb.TransId,
	}
	return strings.Join(pairs, "&")
}

// Helpers to construct create payload
func NewMomoCreatePayload(cfg app.Config, orderID string, amount int64, info string, requestID string) model.MomoCreateRequest {
	return model.MomoCreateRequest{
		PartnerCode: cfg.PartnerCode,
		AccessKey:   cfg.AccessKey,
		RequestId:   requestID,
		Amount:      strconv.FormatInt(amount, 10),
		OrderId:     orderID,
		OrderInfo:   info,
		RedirectUrl: cfg.AppBaseURL + "/momo/return",
		IpnUrl:      cfg.AppBaseURL + "/api/payment/momo/ipn",
		//Lang:        cfg.Lang,
		ExtraData:   "",
		RequestType: "captureWallet",
	}
}

func NewRequestID() string { return fmt.Sprintf("req_%d", time.Now().UnixMilli()) }
