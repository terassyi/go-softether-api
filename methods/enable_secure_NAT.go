package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type EnableSecureNAT struct {
	pkg.Base
	Params *EnableSecureNATParams `json:"params"`
}

func (g *EnableSecureNAT) Parameter() pkg.Params {
	return g.Params
}

func NewEnableSecureNAT(hubName string) *EnableSecureNAT {
	return &EnableSecureNAT{
		Base:   pkg.NewBase("EnableSecureNAT"),
		Params: newEnableSecureNATParams(hubName),
	}
}

func (g *EnableSecureNAT) Name() string {
	return g.Base.Name
}

func (g *EnableSecureNAT) GetId() int {
	return g.Id
}

func (g *EnableSecureNAT) SetId(id int) {
	g.Base.Id = id
}

func (g *EnableSecureNAT) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnableSecureNATParams struct {
	HubName string `json:"HubName_str"`
}

func newEnableSecureNATParams(hubName string) *EnableSecureNATParams {
	return &EnableSecureNATParams{
		HubName: hubName,
	}
}

func (p *EnableSecureNATParams) Tags() []string {
	return []string{"HubName_str"}
}
