package methods

import (
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type SetHub struct {
	pkg.Base
	Params *SetHubParams `json:"params"`
}

func NewSetHub(name, password string, online bool) *SetHub {
	return &SetHub{
		Base:   pkg.NewBase("SetHub"),
		Params: newSetHubParams(name, password, online),
	}
}

func (m *SetHub) Name() string {
	return m.Base.Name
}

func (m *SetHub) GetId() int {
	return m.Id
}

func (m *SetHub) SetId(id int) {
	m.Base.Id = id
}

func (m *SetHub) Parameter() pkg.Params {
	return m.Params
}

func (m *SetHub) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type SetHubParams struct {
	HubName                string `json:"HubName_str"`
	AdminPasswordPlainText string `json:"AdminPasswordPlainText_str"`
	Online                 bool   `json:"Online_bool"`
	MaxSession             int    `json:"MaxSession_u32"`
	NoEnum                 bool   `json:"NoEnum_bool"`
	HubType                int    `json:"HubType_u32"`
}

func newSetHubParams(name, password string, online bool) *SetHubParams {
	return &SetHubParams{
		HubName:                name,
		AdminPasswordPlainText: password,
		Online:                 online,
		MaxSession:             0,
		NoEnum:                 false,
		HubType:                0,
	}
}

func (p *SetHubParams) Tags() []string {
	return []string{
		"HubName_str",
		"AdminPasswordPlainText_str",
		"Online_bool",
		"MaxSession_u32",
		"NoEnum_bool",
		"HubType_u32",
	}
}
