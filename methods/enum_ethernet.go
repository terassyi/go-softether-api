package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type EnumEthernet struct {
	pkg.Base
	Params *EnumEthernetParams `json:"params"`
}

func (g *EnumEthernet) Parameter() pkg.Params {
	return g.Params
}

func NewEnumEthernet() *EnumEthernet {
	return &EnumEthernet{
		Base:   pkg.NewBase("EnumEthernet"),
		Params: newEnumEthernetParams(),
	}
}

func (g *EnumEthernet) Name() string {
	return g.Base.Name
}

func (g *EnumEthernet) GetId() int {
	return g.Id
}

func (g *EnumEthernet) SetId(id int) {
	g.Base.Id = id
}

func (g *EnumEthernet) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnumEthernetParams struct {
}

func newEnumEthernetParams() *EnumEthernetParams {
	return &EnumEthernetParams{}
}

func (p *EnumEthernetParams) Tags() []string {
	return []string{}
}
