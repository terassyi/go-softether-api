package methods

import (
	"bytes"
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type GetHub struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func NewGetHub(name string) *GetHub {
	return &GetHub{
		Base:   pkg.NewBase("GetHub"),
		Params: newGetHubParams(name),
	}
}

func (m *GetHub) Name() string {
	return m.Base.Name
}

func (m *GetHub) GetId() int {
	return m.Id
}

func (m *GetHub) SetId(id int) {
	m.Base.Id = id
}

func (m *GetHub) Parameter() pkg.Params {
	return m.Params
}

func (m *GetHub) Marshall() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type GetHubParams struct {
	HubName string `json:"HubName_str"`
}

func newGetHubParams(name string) *GetHubParams {
	return &GetHubParams{HubName: name}
}

func (p *GetHubParams) Tags() []string {
	return []string{
		"HubName_str",
	}
}
