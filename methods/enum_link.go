package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type EnumLink struct {
	pkg.Base
	Params *EnumLinkParams `json:"params"`
}

func (g *EnumLink) Parameter() pkg.Params {
	return g.Params
}

func NewEnumLink(name string) *EnumLink {
	return &EnumLink{
		Base:   pkg.NewBase("EnumLink"),
		Params: newEnumLinkParams(name),
	}
}

func (g *EnumLink) Name() string {
	return g.Base.Name
}

func (g *EnumLink) GetId() int {
	return g.Id
}

func (g *EnumLink) SetId(id int) {
	g.Base.Id = id
}

func (g *EnumLink) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnumLinkParams struct {
	HubName string `json:"HubName_str"`
}

func newEnumLinkParams(name string) *EnumLinkParams {
	return &EnumLinkParams{HubName: name}
}

func (p *EnumLinkParams) Tags() []string {
	return []string{
		"HubName_str",
	}
}
