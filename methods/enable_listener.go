package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/terassyi/go-softether-api/pkg"
)

type EnableListener struct {
	pkg.Base
	Params pkg.Params `josn:"params"`
}

func (g *EnableListener) Parameter() pkg.Params {
	return g.Params
}

func NewEnableListener(port int, enable bool) *EnableListener {
	return &EnableListener{
		Base:   pkg.NewBase("EnableListener"),
		Params: newEnableListenerParams(port, enable),
	}
}

func (g *EnableListener) Name() string {
	return g.Base.Name
}

func (g *EnableListener) GetId() int {
	return g.Id
}

func (g *EnableListener) SetId(id int) {
	g.Base.Id = id
}

func (g *EnableListener) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnableListenerParams struct {
	Port   int  `json:"Port_u32"`
	Enable bool `json:"Enable_bool"`
}

func newEnableListenerParams(port int, enable bool) *EnableListenerParams {
	return &EnableListenerParams{
		Port:   port,
		Enable: enable,
	}
}

func (p *EnableListenerParams) Set(key string, val interface{}) error {
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

func (p *EnableListenerParams) Get(key string) (interface{}, error) {
	switch key {
	case "Port":
		return p.Port, nil
	case "Enable":
		return p.Enable, nil
	default:
		return nil, fmt.Errorf("not found such a parameter")
	}
}
