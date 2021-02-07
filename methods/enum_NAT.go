package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type EnumNAT struct {
	pkg.Base
	Params *EnumNATParams `json:"params"`
}

func (g *EnumNAT) Parameter() pkg.Params {
	return g.Params
}

func NewEnumNAT(hubName string) *EnumNAT {
	return &EnumNAT{
		Base:   pkg.NewBase("EnumNAT"),
		Params: newEnumNATParams(hubName),
	}
}

func (g *EnumNAT) Name() string {
	return g.Base.Name
}

func (g *EnumNAT) GetId() int {
	return g.Id
}

func (g *EnumNAT) SetId(id int) {
	g.Base.Id = id
}

func (g *EnumNAT) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnumNATParams struct {
	HubName string `json:"HubName_str"`
}

func newEnumNATParams(hubName string) *EnumNATParams {
	return &EnumNATParams{
		HubName: hubName,
	}
}

func (p *EnumNATParams) Tags() []string {
	return []string{"HubName_str"}
}
