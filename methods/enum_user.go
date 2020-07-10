package methods

import (
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type EnumUser struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func (g *EnumUser) Parameter() pkg.Params {
	return g.Params
}

func NewEnumUser(hub string) *EnumUser {
	return &EnumUser{
		Base:   pkg.NewBase("EnumUser"),
		Params: newEnumUserParams(hub),
	}
}

func (g *EnumUser) Name() string {
	return g.Base.Name
}

func (g *EnumUser) GetId() int {
	return g.Id
}

func (g *EnumUser) SetId(id int) {
	g.Base.Id = id
}

func (g *EnumUser) Marshall() ([]byte, error) {
	return json.Marshal(g)
}

type EnumUserParams struct {
	HubName string `json:"HubName_str"`
}

func newEnumUserParams(hub string) *EnumUserParams {
	return &EnumUserParams{
		HubName: hub,
	}
}

func (p *EnumUserParams) Tags() []string {
	return []string{"HubName_str"}
}
