package dana

import "time"

type ResponseBody struct {
	Response  Response `json:"response" valid:"required"`
	Signature string   `json:"signature" valid:"required"`
}

type Response struct {
	Head ResponseHeader `json:"head" valid:"required"`
	Body interface{}    `json:"body" valid:"required"`
}

type ResponseHeader struct {
	Function  string `json:"function" valid:"required"`
	ClientID  string `json:"clientId" valid:"required"`
	Version   string `json:"version" valid:"required"`
	RespTime  string `json:"respTime" valid:"required"`
	RespMsgID string `json:"reqMsgId" valid:"required"`
}

type OrderResponseData struct {
	MerchantTransID *string    `json:"merchantTransId,omitempty" valid:"optional"`
	AcquirementID   *string    `json:"acquirementId,omitempty" valid:"optional"`
	CheckoutURL     *string    `json:"checkoutUrl,omitempty" valid:"optional"`
	ResultInfo      ResultInfo `json:"resultInfo" valid:"required"`
}

type OrderDetailData struct {
	ResultInfo      ResultInfo     `json:"resultInfo" valid:"required"`
	AcquirementID   string         `json:"acquirementId" valid:"optional"`
	MerchantTransID string         `json:"merchantTransId" valid:"optional"`
	Buyer           InputUserInfo  `json:"buyer" valid:"optional"`
	Seller          InputUserInfo  `json:"seller" valid:"optional"`
	OrderTitle      string         `json:"orderTitle" valid:"optional"`
	ExtendedInfo    string         `json:"extendedInfo" valid:"optional"`
	AmountDetail    AmountDetail   `json:"amountDetail" valid:"optional"`
	TimeDetail      TimeDetail     `json:"timeDetail" valid:"optional"`
	StatusDetail    StatusDetail   `json:"statusDetail" valid:"optional"`
	Goods           []Good         `json:"goods" valid:"optional"`
	ShippingInfo    []ShippingInfo `json:"shippingInfo" valid:"optional"`
	OrderMemo       string         `json:"orderMemo" valid:"optional"`
	PaymentViews    []PaymentView  `json:"paymentViews" valid:"optional"`
}

type RefundResponseData struct {
	ResultInfo ResultInfo `json:"resultInfo" valid:"required"`
	RefundID   *string    `json:"refundId,omitempty" valid:"optional"`
	RequestID  *string    `json:"requestId,omitempty" valid:"optional"`
}

type ResultInfo struct {
	ResultStatus string `json:"resultStatus" valid:"optional"`
	ResultCodeID string `json:"resultCodeId" valid:"optional"`
	ResultMsg    string `json:"resultMsg" valid:"optional"`
	ResultCode   string `json:"resultCode" valid:"optional"`
}

type PayFinishResponse struct {
	Response  ResponsePayFinish `json:"response"`
	Signature string            `json:"signature"`
}

type ResponsePayFinish struct {
	Head ResponseHeader        `json:"head"`
	Body ResponseBodyPayFinish `json:"body"`
}

type ResponseBodyPayFinish struct {
	ResultInfo ResultInfo `json:"resultInfo"`
}

type InputUserInfo struct {
	UserID           string `json:"userId" valid:"optional"`
	ExternalUserID   string `json:"externalUserId" valid:"optional"`
	ExternalUserType string `json:"externalUserType" valid:"optional"`
	Nickname         string `json:"nickname" valid:"optional"`
}

type AmountDetail struct {
	OrderAmount      Amount `json:"orderAmount" valid:"required"`
	PayAmount        Amount `json:"payAmount" valid:"optional"`
	VoidAmount       Amount `json:"voidAmount" valid:"optional"`
	ConfirmAmount    Amount `json:"confirmAmount" valid:"optional"`
	RefundAmount     Amount `json:"refundAmount" valid:"optional"`
	ChargebackAmount Amount `json:"chargebackAmount" valid:"optional"`
	ChargeAmount     Amount `json:"chargeAmount" valid:"optional"`
}

type TimeDetail struct {
	CreatedTime    time.Time   `json:"createdTime" valid:"required"`
	ExpiryTime     time.Time   `json:"expiryTime" valid:"required"`
	PaidTimes      []time.Time `json:"paidTimes" valid:"optional"`
	ConfirmedTimes []time.Time `json:"confirmedTimes" valid:"optional"`
	CancelledTime  time.Time   `json:"cancelledTime" valid:"optional"`
}

type StatusDetail struct {
	AcquirementStatus StatusDetailEnum `json:"acquirementStatus" valid:"required"`
	Frozen            bool             `json:"frozen" valid:"required"`
}

type PaymentView struct {
	CashierRequestID     string          `json:"cashierRequestId" valid:"required"`
	PaidTime             time.Time       `json:"paidTime" valid:"required"`
	PayOptionInfos       []PayOptionInfo `json:"payOptionInfos" valid:"required"`
	PayRequestExtendInfo string          `json:"payRequestExtendInfo" valid:"optional"`
	ExtendInfo           string          `json:"extendInfo" valid:"optional"`
}

type PayOptionInfo struct {
	PayMethod               PayMethodEnum `json:"payMethod" valid:"required"`
	PayAmount               Amount        `json:"payAmount" valid:"required"`
	TransAmount             Amount        `json:"transAmount" valid:"optional"`
	ChargeAmount            Amount        `json:"chargeAmount" valid:"optional"`
	ExtendInfo              string        `json:"extendInfo" valid:"optional"`
	PayOptionBillExtendInfo string        `json:"payOptionBillExtendInfo" valid:"optional"`
}
