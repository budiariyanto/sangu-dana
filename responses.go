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
	Version   string `json:"version" valid:"required"`
	Function  string `json:"function" valid:"required"`
	ClientId  string `json:"clientId" valid:"required"`
	RespTime  string `json:"respTime" valid:"required"`
	RespMsgId string `json:"reqMsgId" valid:"required"`
}

type ResponseBody struct {
	MerchantTransId string     `json:"merchantTransId" valid:"optional"`
	AcquirementId   string     `json:"acquirementId" valid:"optional"`
	CheckoutUrl     string     `json:"checkoutUrl" valid:"optional"`
	ResultInfo      ResultInfo `json:"resultInfo" valid:"required"`
}

type ResultInfo struct {
	ResultStatus string `json:"resultStatus" valid:"optional"`
	ResultCodeId string `json:"resultCodeId" valid:"optional"`
	ResultMsg    string `json:"resultMsg" valid:"optional"`
	ResultCode   string `json:"resultCode" valid:"optional"`
}
