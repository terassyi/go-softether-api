package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type EnumLocalBridge struct {
	pkg.Base
	Params *EnumLocalBridgeParams `json:"params"`
}

func (g *EnumLocalBridge) Parameter() pkg.Params {
	return g.Params
}

func NewEnumLocalBridge() *EnumLocalBridge {
	return &EnumLocalBridge{
		Base:   pkg.NewBase("EnumLocalBridge"),
		Params: newEnumLocalBridgeParams(),
	}
}

func (g *EnumLocalBridge) Name() string {
	return g.Base.Name
}

func (g *EnumLocalBridge) GetId() int {
	return g.Id
}

func (g *EnumLocalBridge) SetId(id int) {
	g.Base.Id = id
}

func (g *EnumLocalBridge) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnumLocalBridgeParams struct {
}

func newEnumLocalBridgeParams() *EnumLocalBridgeParams {
	return &EnumLocalBridgeParams{}
}

func (p *EnumLocalBridgeParams) Tags() []string {
	return []string{}
}
