package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/terassyi/go-softether-api/pkg"
)

type CreateListener struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func (g *CreateListener) Parameter() pkg.Params {
	return g.Params
}

func NewCreateListener(port int, enable bool) *CreateListener {
	return &CreateListener{
		Base:   pkg.NewBase("CreateListener"),
		Params: newCreateListenerParams(port, enable),
	}
}

func (g *CreateListener) Name() string {
	return g.Base.Name
}

func (g *CreateListener) GetId() int {
	return g.Id
}

func (g *CreateListener) SetId(id int) {
	g.Base.Id = id
}

func (g *CreateListener) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type CreateListenerParams struct {
	Port   int  `json:"Port_u32"`
	Enable bool `json:"Enable_bool"`
}

func newCreateListenerParams(port int, enable bool) *CreateListenerParams {
	return &CreateListenerParams{
		Port:   port,
		Enable: enable,
	}
}

func (p *CreateListenerParams) Set(key string, val interface{}) error {
	switch key {
	case "Port":
		switch v := val.(type) {
		case int:
			{
				p.Port = v
				return nil
			}
		default:
			return fmt.Errorf("invalid parameter type.")
		}
	case "Enable":
		switch v := val.(type) {
		case bool:
			{
				p.Enable = v
				return nil
			}
		default:
			return fmt.Errorf("invalid parameter type")
		}
	}
	return fmt.Errorf("not found such a parameter")
}

func (p *CreateListenerParams) Get(key string) (interface{}, error) {
	switch key {
	case "Port":
		return p.Port, nil
	case "Enable":
		return p.Enable, nil
	default:
		return nil, fmt.Errorf("not found such a parameter")
	}
}
