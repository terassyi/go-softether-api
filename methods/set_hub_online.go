package methods

import (
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type SetHubOnline struct {
	pkg.Base
	Params *SetHubOnlineParams `json:"params"`
}

func NewSetHubOnline(name string, online bool) *SetHubOnline {
	return &SetHubOnline{
		Base:   pkg.NewBase("SetHubOnline"),
		Params: newSetHubOnlineParams(name, online),
	}
}

func (m *SetHubOnline) Name() string {
	return m.Base.Name
}

func (m *SetHubOnline) GetId() int {
	return m.Id
}

func (m *SetHubOnline) SetId(id int) {
	m.Base.Id = id
}

func (m *SetHubOnline) Parameter() pkg.Params {
	return m.Params
}

func (m *SetHubOnline) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type SetHubOnlineParams struct {
	HubName string `json:"HubName_str"`
	Online  bool   `json:"Online_bool"`
}

func newSetHubOnlineParams(name string, online bool) *SetHubOnlineParams {
	return &SetHubOnlineParams{
		HubName: name,
		Online:  online,
	}
}

func (p *SetHubOnlineParams) Tags() []string {
	return []string{
		"HubName_str",
		"Online_bool",
	}
}
