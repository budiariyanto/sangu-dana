package dana

type OrderResponse struct {
	Response  Response `json:"response" valid:"required"`
	Signature string   `json:"signature" valid:"required"`
}
type Response struct {
	Head ResponseHeader `json:"head" valid:"required"`
	Body ResponseBody   `json:"body" valid:"required"`
}

type ResponseHeader struct {
	Function  string `json:"function" valid:"required"`
	ClientID  string `json:"clientId" valid:"required"`
	Version   string `json:"version" valid:"required"`
	RespTime  string `json:"respTime" valid:"required"`
	RespMsgID string `json:"reqMsgId" valid:"required"`
}

type ResponseBody struct {
	MerchantTransID string     `json:"merchantTransId" valid:"optional"`
	AcquirementID   string     `json:"acquirementId" valid:"optional"`
	CheckoutURL     string     `json:"checkoutUrl" valid:"optional"`
	ResultInfo      ResultInfo `json:"resultInfo" valid:"required"`
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
