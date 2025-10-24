package model

type CreatePaymentRequest struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

type CreatePaymentResponse struct {
	OrderID   string `json:"orderId"`
	RequestID string `json:"requestId"`
	PayURL    string `json:"payUrl"`
	Message   string `json:"message"`
	Code      int    `json:"resultCode"`
}

// MoMo v2 create
// Note: amount must be string in payload

type MomoCreateRequest struct {
	PartnerCode string `json:"partnerCode"`
	AccessKey   string `json:"accessKey"`
	RequestId   string `json:"requestId"`
	Amount      string `json:"amount"`
	OrderId     string `json:"orderId"`
	OrderInfo   string `json:"orderInfo"`
	RedirectUrl string `json:"redirectUrl"`
	IpnUrl      string `json:"ipnUrl"`
	//Lang        string `json:"lang,omitempty"`
	ExtraData   string `json:"extraData"`
	RequestType string `json:"requestType"`
	Signature   string `json:"signature"`
}

type MomoCreateResponse struct {
	PartnerCode  string `json:"partnerCode"`
	OrderId      string `json:"orderId"`
	RequestId    string `json:"requestId"`
	Amount       int64  `json:"amount"`
	ResponseTime int64  `json:"responseTime"`
	Message      string `json:"message"`
	ResultCode   int    `json:"resultCode"`
	PayUrl       string `json:"payUrl"`
	Signature    string `json:"signature"`
}

// Callback (IPN/Return)

type MomoCallback struct {
	PartnerCode  string `json:"partnerCode"`
	OrderId      string `json:"orderId"`
	RequestId    string `json:"requestId"`
	Amount       string `json:"amount"`
	OrderInfo    string `json:"orderInfo"`
	OrderType    string `json:"orderType"`
	TransId      string `json:"transId"`
	ResultCode   string `json:"resultCode"`
	Message      string `json:"message"`
	PayType      string `json:"payType"`
	ResponseTime string `json:"responseTime"`
	ExtraData    string `json:"extraData"`
	Signature    string `json:"signature"`
}

type QueryStatusRequest struct {
	PartnerCode string `json:"partnerCode"`
	AccessKey   string `json:"accessKey"`
	RequestId   string `json:"requestId"`
	OrderId     string `json:"orderId"`
	Lang        string `json:"lang,omitempty"`
	Signature   string `json:"signature"`
}

type QueryStatusResponse struct {
	PartnerCode  string `json:"partnerCode"`
	OrderId      string `json:"orderId"`
	RequestId    string `json:"requestId"`
	Amount       int64  `json:"amount"`
	ResponseTime int64  `json:"responseTime"`
	Message      string `json:"message"`
	ResultCode   int    `json:"resultCode"`
	TransId      int64  `json:"transId"`
	PayType      string `json:"payType"`
	Signature    string `json:"signature"`
}
