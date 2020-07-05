package methods

import (
	"encoding/json"
	"fmt"
	"github.com/terassyi/go-softether-api/pkg"
)

type Test struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func NewTest() *Test {
	return &Test{
		Base:   pkg.NewBase("Test"),
		Params: &TestParams{IntValue: 0},
	}
}

func (t *Test) Name() string {
	return t.Base.Name
}

func (t *Test) GetId() int {
	return t.Id
}

func (t *Test) SetId(id int) {
	t.Base.Id = id
}

func (t *Test) Parameter() pkg.Params {
	return t.Params
}

func (t *Test) Marshall() ([]byte, error) {
	return json.Marshal(t)
}

type TestParams struct {
	IntValue int `json:"IntValue_u32"`
}

func (p *TestParams) Set(key string, val interface{}) error {
	switch key {
	case "IntValue":
		{
			switch v := val.(type) {
			case int:
				{
					p.IntValue = v
					return nil
				}
			default:
				return fmt.Errorf("invalid parameter type")
			}
		}
	}
	return fmt.Errorf("not found such a parameter")
}

func (p *TestParams) Get(key string) (interface{}, error) {
	switch key {
	case "IntValue":
		return p.IntValue, nil
	default:
		return nil, fmt.Errorf("not found such a parameter")
	}
}
