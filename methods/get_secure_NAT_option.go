package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type GetSecureNATOption struct {
	pkg.Base
	Params *GetSecureNATOptionParams `json:"params"`
}

func (g *GetSecureNATOption) Parameter() pkg.Params {
	return g.Params
}

func NewGetSecureNATOption(rpcHubName string) *GetSecureNATOption {
	return &GetSecureNATOption{
		Base:   pkg.NewBase("GetSecureNATOption"),
		Params: newGetSecureNATOptionParams(rpcHubName),
	}
}

func (g *GetSecureNATOption) Name() string {
	return g.Base.Name
}

func (g *GetSecureNATOption) GetId() int {
	return g.Id
}

func (g *GetSecureNATOption) SetId(id int) {
	g.Base.Id = id
}

func (g *GetSecureNATOption) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type GetSecureNATOptionParams struct {
	RpcHubName string `json:"RpcHubName_str"`
}

func newGetSecureNATOptionParams(rpcHubName string) *GetSecureNATOptionParams {
	return &GetSecureNATOptionParams{
		RpcHubName: rpcHubName,
	}
}

func (p *GetSecureNATOptionParams) Tags() []string {
	return []string{"RpcHubName_str"}
}
