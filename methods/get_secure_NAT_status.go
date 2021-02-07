package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type GetSecureNATStatus struct {
	pkg.Base
	Params *GetSecureNATStatusParams `json:"params"`
}

func (g *GetSecureNATStatus) Parameter() pkg.Params {
	return g.Params
}

func NewGetSecureNATStatus(hubName string) *GetSecureNATStatus {
	return &GetSecureNATStatus{
		Base:   pkg.NewBase("GetSecureNATStatus"),
		Params: newGetSecureNATStatusParams(hubName),
	}
}

func (g *GetSecureNATStatus) Name() string {
	return g.Base.Name
}

func (g *GetSecureNATStatus) GetId() int {
	return g.Id
}

func (g *GetSecureNATStatus) SetId(id int) {
	g.Base.Id = id
}

func (g *GetSecureNATStatus) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type GetSecureNATStatusParams struct {
	HubName string `json:"HubName_str"`
}

func newGetSecureNATStatusParams(hubName string) *GetSecureNATStatusParams {
	return &GetSecureNATStatusParams{
		HubName: hubName,
	}
}

func (p *GetSecureNATStatusParams) Tags() []string {
	return []string{"HubName_str"}
}
