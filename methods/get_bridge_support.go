package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type GetBridgeSupport struct {
	pkg.Base
	Params *GetBridgeSupportParams `json:"params"`
}

func (g *GetBridgeSupport) Parameter() pkg.Params {
	return g.Params
}

func NewGetBridgeSupport() *GetBridgeSupport {
	return &GetBridgeSupport{
		Base:   pkg.NewBase("GetBridgeSupport"),
		Params: newGetBridgeSupportParams(),
	}
}

func (g *GetBridgeSupport) Name() string {
	return g.Base.Name
}

func (g *GetBridgeSupport) GetId() int {
	return g.Id
}

func (g *GetBridgeSupport) SetId(id int) {
	g.Base.Id = id
}

func (g *GetBridgeSupport) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type GetBridgeSupportParams struct {
}

func newGetBridgeSupportParams() *GetBridgeSupportParams {
	return &GetBridgeSupportParams{}
}

func (p *GetBridgeSupportParams) Tags() []string {
	return []string{}
}
