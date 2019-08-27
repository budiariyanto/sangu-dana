package dana

type OrderRequest struct {
	Request   Request `json:"request" valid:"required"`
	Signature string  `json:"signature" valid:"required"`
}

type Request struct {
	Head RequestHeader `json:"head" valid:"required"`
	Body RequestBody   `json:"body" valid:"required"`
}

type RequestHeader struct {
	Version      string `json:"version" valid:"required"`
	Function     string `json:"function" valid:"required"`
	ClientID     string `json:"clientId" valid:"required"`
	ReqTime      string `json:"reqTime" valid:"required"`
	ReqMsgID     string `json:"reqMsgId" valid:"required"`
	ClientSecret string `json:"clientSecret" valid:"required"`
	AccessToken  string `json:"accessToken" valid:"optional"`
	Reserve      string `json:"reserve" valid:"optional"`
}

type RequestBody struct {
	Order             Order             `json:"order" valid:"required"`
	MerchantID        string            `json:"merchantId" valid:"required"`
	Mcc               string            `json:"mcc" valid:"optional"`
	ProductCode       string            `json:"productCode" valid:"required"`
	EnvInfo           EnvInfo           `json:"envInfo" valid:"required"`
	NotificationUrls  []NotificationUrl `json:"notificationUrls" valid:"optional"`
	ExtendInfo        string            `json:"extendInfo" valid:"optional"`
	PaymentPreference PaymentPreference `json:"paymentPreference" valid:"optional"`
}

type Order struct {
	OrderTitle        string         `json:"orderTitle"`
	OrderAmount       Amount         `json:"orderAmount"`
	MerchantTransID   string         `json:"merchantTransId"`
	MerchantTransType string         `json:"merchantTransType"`
	OrderMemo         string         `json:"orderMemo"`
	CreatedTime       string         `json:"createdTime"`
	ExpiryTime        string         `json:"expiryTime"`
	Goods             []Good         `json:"goods"`
	ShippingInfo      []ShippingInfo `json:"shippingInfo"`
}

type Amount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type Good struct {
	MerchantGoodsID    string `json:"merchantGoodsId"`
	Description        string `json:"description"`
	Category           string `json:"category"`
	Price              Amount `json:"price"`
	Unit               string `json:"unit"`
	Quantity           string `json:"quantity"`
	MerchantShippingID string `json:"merchantShippingId"`
	SnapshotURL        string `json:"snapshotUrl"`
	ExtendInfo         string `json:"extendInfo"`
}

type ShippingInfo struct {
	MerchantShippingID string `json:"merchantShippingId"`
	TrackingNo         string `json:"trackingNo"`
	Carrier            string `json:"carrier"`
	ChargeAmount       Amount `json:"chargeAmount"`
	CountryName        string `json:"countryName"`
	StateName          string `json:"stateName"`
	CityName           string `json:"cityName"`
	AreaName           string `json:"areaName"`
	Address1           string `json:"address1"`
	Address2           string `json:"address2"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	MobileNo           string `json:"mobileNo"`
	PhoneNo            string `json:"phoneNo"`
	ZipCode            string `json:"zipCode"`
	Email              string `json:"email"`
	FaxNo              string `json:"faxNo"`
}

type EnvInfo struct {
	SessionID          string `json:"sessionId"`
	TokenID            string `json:"tokenId"`
	WebsiteLanguage    string `json:"websiteLanguage"`
	ClientIP           string `json:"clientIp"`
	OsType             string `json:"osType"`
	AppVersion         string `json:"appVersion"`
	SdkVersion         string `json:"sdkVersion"`
	SourcePlatform     string `json:"sourcePlatform"`
	TerminalType       string `json:"terminalType"`
	ClientKey          string `json:"clientKey"`
	OrderTerminalType  string `json:"orderTerminalType"`
	OrderOsType        string `json:"orderOsType"`
	MerchantAppVersion string `json:"merchantAppVersion"`
	ExtendInfo         string `json:"extendInfo"`
}

type NotificationUrl struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

type PaymentPreference struct {
	DisabledPayMethods string `json:"disabledPayMethods"`
}

type PayFinishRequest struct {
	Request   RequestPayFinish `json:"request"`
	Signature string           `json:"signature"`
}
type RequestPayFinish struct {
	Head RequestHeader        `json:"head"`
	Body RequestBodyPayFinish `json:"body"`
}

type RequestBodyPayFinish struct {
	AcquirementID     string `json:"acquirementId"`
	MerchantTransID   string `json:"merchantTransId"`
	FinishedTime      string `json:"finishedTime"`
	CreatedTime       string `json:"createdTime"`
	MerchantID        string `json:"merchantId"`
	OrderAmount       Amount `json:"orderAmount"`
	AcquirementStatus string `json:"acquirementStatus"`
	ExtendInfo        string `json:"extendInfo"`
}
