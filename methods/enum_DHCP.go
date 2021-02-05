package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type EnumDHCP struct {
	pkg.Base
	Params *EnumDHCPParams `json:"params"`
}

func (g *EnumDHCP) Parameter() pkg.Params {
	return g.Params
}

func NewEnumDHCP(hubName string) *EnumDHCP {
	return &EnumDHCP{
		Base:   pkg.NewBase("EnumDHCP"),
		Params: newEnumDHCPParams(hubName),
	}
}

func (g *EnumDHCP) Name() string {
	return g.Base.Name
}

func (g *EnumDHCP) GetId() int {
	return g.Id
}

func (g *EnumDHCP) SetId(id int) {
	g.Base.Id = id
}

func (g *EnumDHCP) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnumDHCPParams struct {
	HubName string `json:"HubName_str"`
}

func newEnumDHCPParams(hubName string) *EnumDHCPParams {
	return &EnumDHCPParams{
		HubName: hubName,
	}
}

func (p *EnumDHCPParams) Tags() []string {
	return []string{"HubName_str"}
}
