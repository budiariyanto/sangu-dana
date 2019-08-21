package dana

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"strings"
	"time"
)

const (
	//ORDER_PATH       = "https://api-sandbox.saas.dana.id/alipayplus/acquiring/order/createOrder.htm"
	ORDER_PATH       = "alipayplus/acquiring/order/createOrder.htm"
	QUERY_PATH       = "alipayplus/acquiring/order/query.htm"
	DANA_TIME_LAYOUT = "2006-01-02T15:04:05.000-07:00"
	CURRENCY_IDR     = "IDR"
)

// CoreGateway struct
type CoreGateway struct {
	Client Client
}

// Call : base method to call Core API
func (gateway *CoreGateway) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.BaseUrl + path

	return gateway.Client.Call(method, path, header, body, v)
}

func (gateway *CoreGateway) Order(reqBody RequestBody) (res OrderResponse, err error) {
	now := time.Now()

	head := RequestHeader{}
	head.Version = gateway.Client.Version
	head.Function = gateway.Client.Function
	head.ClientId = gateway.Client.ClientId
	head.ReqTime = now.Format(DANA_TIME_LAYOUT)
	head.ReqMsgId = uuid.Must(uuid.NewV4()).String()
	head.ClientSecret = gateway.Client.ClientSecret

	req := Request{
		Head: head,
		Body: reqBody,
	}

	sig, err := generateSignature(req)
	if err != nil {
		err = fmt.Errorf("failed to generate signature: %v", err)
		return
	}
	orderReq := OrderRequest{
		Request:   req,
		Signature: sig,
	}

	reqJson, _ := json.Marshal(orderReq)
	log.Println("Dana request: ", string(reqJson))
	requestByte, _ := json.Marshal(orderReq)
	requestReader := bytes.NewBuffer(requestByte)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	err = gateway.Call("POST", ORDER_PATH, headers, requestReader, &res)
	if err != nil {
		return
	}

	//response RSA verification
	err = verifyResponse(res.Response, res.Signature)
	if err != nil {
		log.Printf("could not unsign request: %v", err)
	}
	return
}

func verifyResponse(res Response, sig string) error {
	parser, perr := loadPublicKey()
	if perr != nil {
		log.Printf("could load public key: %v", perr)
	}
	resp, _ := json.Marshal(res)
	ds, _ := base64.StdEncoding.DecodeString(sig)
	return parser.Unsign(resp, ds)
}
