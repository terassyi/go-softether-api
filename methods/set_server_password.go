package methods

import (
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type SetServerPassword struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func NewSetServerPassword() *SetServerPassword {
	return &SetServerPassword{
		Base:   pkg.NewBase("SetServerPassword"),
		Params: newSetServerPasswordParams("test"),
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

func (p *SetServerPasswordParams) Tags() []string {
	return []string{"PlainTextPassword_str"}
}
