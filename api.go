package softetherapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/terassyi/go-softether-api/pkg"
	"io/ioutil"
	"net/http"
)

//to handle softether vpn server json-rpc api
type Api struct {
	Host     string
	Port     int
	Hub      string
	Password string
	id       Id
	conn     *http.Client
}

// Id interface is used for request id
type Id interface {
	Incl()
	Describe() int
}

// id struct is the implement of Id interface
type id int

func initId() Id {
	return id(0)
}

// This method increment an id
func (i id) Incl() {
	i += 1
}

func (i id) Describe() int {
	return int(i)
}

func New(host string, port int, hub, password string) *Api {
	return &Api{
		Host:     host,
		Port:     port,
		Hub:      hub,
		Password: password,
		id:       initId(),
		conn:     connect(),
	}
}

func (api *Api) entrypoint() string {
	return fmt.Sprintf("https://%s:%d/api", api.Host, api.Port)
}

func (api *Api) Call(method pkg.Method) (map[string]interface{}, error) {
	api.id.Incl()
	method.SetId(api.id.Describe())
	body, err := method.Marshall()
	fmt.Println(string(body))
	if err != nil {
		return nil, fmt.Errorf("[error] failed to marshall request: %v", err)
	}
	req, err := http.NewRequest("POST", api.entrypoint(), bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("[error] failed to create new http request: %v", err)
	}
	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-VPNADMIN-HUBNAME", api.Hub)
	req.Header.Set("X-VPNADMIN-PASSWORD", api.Password)
	// send an api request
	res, err := api.conn.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[error] failed to call: %v", err)
	}
	return validateResponse(res)
}

func connect() *http.Client {
	tl := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return &http.Client{Transport: tl}
}

func validateResponse(res *http.Response) (map[string]interface{}, error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("[error] failed to read: %v", err)
	}
	fmt.Println(string(body))
	b := make(map[string]interface{})
	err = json.Unmarshal(body, &b)
	if err != nil {
		return nil, err
	}
	if _, ok := b["result"]; ok {
		return b["result"].(map[string]interface{}), nil
	}
	return nil, fmt.Errorf("%v", b["error"])
}
