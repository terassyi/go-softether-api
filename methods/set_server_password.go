package methods

import (
	"encoding/json"
	"fmt"
	"github.com/terassyi/go-softether-api/pkg"
)

type SetServerPassword struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func NewSetServerPassword() *SetServerPassword {
	return &SetServerPassword{
		Base:   pkg.NewBase("SetServerPassword"),
		Params: &SetServerPasswordParams{IntValue: 0},
	}
}

func (m *SetServerPassword) Name() string {
	return m.Base.Name
}

func (m *SetServerPassword) GetId() int {
	return m.Id
}

func (m *SetServerPassword) SetId(id int) {
	m.Base.Id = id
}

func (m *SetServerPassword) Parameter() pkg.Params {
	return m.Params
}

func (m *SetServerPassword) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type SetServerPasswordParams struct {
	PlainTextPassword string `json:"PlainTextPassword_str"`
}

func newSetServerPasswordParams(text string) *SetServerPasswordParams {
	return &SetServerPasswordParams{PlainTextPassword: text}
}

func (p *SetServerPasswordParams) Set(key string, val interface{}) error {
	switch key {
	case "PlainTextPassword":
		{
			switch v := val.(type) {
			case string:
				{
					p.PlainTextPassword = v
					return nil
				}
			default:
				return fmt.Errorf("invalid parameter type")
			}
		}
	}
	return fmt.Errorf("not found such a parameter")
}

func (p *SetServerPasswordParams) Get(key string) (interface{}, error) {
	switch key {
	case "PlainTextPassword":
		return p.PlainTextPassword, nil
	default:
		return nil, fmt.Errorf("not found such a parameter")
	}
}
