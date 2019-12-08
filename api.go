package softetherapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//json-rpc const field
const (
	Jsonrpc      = "2.0"
	AuthType_u32 = 1 //password authentication

)

//to handle softether vpn server json-rpc api

type SoftEtherAPIConn struct {
	Host string
	Port string
	// Conn net.Conn
	Conn     *http.Client
	Hub      string
	Password string
	Id       APIId
}

type APIMethod interface {
	Name() string
	GetId() int
	ToJSON() ([]byte, error)
}

type Method struct {
	BaseField
	params Params
}

type BaseField struct {
	Jsonrpc string
	Id      APIId
	Method  string
}

func newBaseField(method string, id APIId) BaseField {
	return BaseField{
		Jsonrpc: Jsonrpc,
		Id:      id,
		Method:  method,
	}
}

func (api *SoftEtherAPIConn) NewRequest(base BaseField, p Params) APIMethod {
	return Method{
		BaseField: base,
		params:    p,
	}
}

type Params interface {
	Show()
	ToMap() map[string]interface{}
}

type APIId int

func (i *APIId) call() {
	*i += 1
}

func NewSoftEtherAPIConn(host string, port int) *SoftEtherAPIConn {
	return &SoftEtherAPIConn{
		Host: host,
		Port: strconv.Itoa(port),
	}
}

func (api *SoftEtherAPIConn) Connect() {
	// sock, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", api.Host, api.Port), &tls.Config{InsecureSkipVerify: true})
	tl := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	sock := &http.Client{
		Transport: tl,
	}
	api.Conn = sock
}

// func (api *SoftEtherAPIConn) Close() error {
// 	return api.Conn.
// }

func (api *SoftEtherAPIConn) Request(method APIMethod) (*http.Response, error) {
	//convert the method to json
	body, err := method.ToJSON()
	if err != nil {
		return nil, err
	}
	//build request body
	req, err := http.NewRequest("POST", api.url(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	//set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-VPNADMIN-HUBNAME", api.Hub)
	req.Header.Set("X-VPNADMIN-PASSWORD", api.Password)
	fmt.Println(*req)
	//send request
	return api.Conn.Do(req)
}

func Response(res *http.Response) (map[string]interface{}, error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//body is json format
	mapedBody := make(map[string]interface{})
	err = json.Unmarshal(body, &mapedBody)
	if err != nil {
		return nil, err
	}
	log.Println(mapedBody)
	return mapedBody, nil
}

func ValidateResponse(res map[string]interface{}) bool {
	if _, ok := res["result"]; ok {
		return ok
	}
	return false
}

func (api *SoftEtherAPIConn) url() string {
	return fmt.Sprintf("https://%s:%s/api", api.Host, api.Port)
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxMask = 0x3F // 63 0b111111
)

func GeneratePassword() string {
	rand.Seed(time.Now().UTC().UnixNano())
	pass := make([]byte, 6)
	for i := range pass {
		pass[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(pass)
}

func (m Method) Name() string {
	return m.Method
}

func (m Method) GetId() int {
	return int(m.Id)
}

func (m Method) ToMap() (map[string]interface{}, error) {
	methodMap := make(map[string]interface{})
	methodMap["jsonrpc"] = m.Jsonrpc
	methodMap["id"] = int(m.Id)
	methodMap["method"] = m.Method

	paramMap := make(map[string]interface{})
	switch method := m.params.(type) {
	case CreateUserParams:
		paramMap = method.ToMap()
		methodMap["params"] = paramMap
		return methodMap, nil
	default:
		paramMap = nil
		return nil, errors.New("Can't decode to map.")
	}
}

func (m Method) ToJSON() ([]byte, error) {
	reqMap, err := m.ToMap()
	if err != nil {
		return nil, err
	}
	return json.Marshal(reqMap)
}
