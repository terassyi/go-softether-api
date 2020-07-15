package methods

import (
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type CreateHub struct {
	pkg.Base
	Params *CreateHubParams `json:"params"`
}

func NewCreateHub(name, password string, online bool) *CreateHub {
	return &CreateHub{
		Base:   pkg.NewBase("CreateHub"),
		Params: newCreateHubParams(name, password, online),
	}
}

func (m *CreateHub) Name() string {
	return m.Base.Name
}

func (m *CreateHub) GetId() int {
	return m.Id
}

func (m *CreateHub) SetId(id int) {
	m.Base.Id = id
}

func (m *CreateHub) Parameter() pkg.Params {
	return m.Params
}

func (m *CreateHub) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type CreateHubParams struct {
	HubName                string `json:"HubName_str"`
	AdminPasswordPlainText string `json:"AdminPasswordPlainText_str"`
	Online                 bool   `json:"Online_bool"`
	MaxSession             int    `json:"MaxSession_u32"`
	NoEnum                 bool   `json:"NoEnum_bool"`
	HubType                int    `json:"HubType_u32"`
}

func newCreateHubParams(name, password string, online bool) *CreateHubParams {
	return &CreateHubParams{
		HubName:                name,
		AdminPasswordPlainText: password,
		Online:                 online,
		MaxSession:             0,
		NoEnum:                 false,
		HubType:                0,
	}
}

func (p *CreateHubParams) Tags() []string {
	return []string{
		"HubName_str",
		"AdminPasswordPlainText_str",
		"Online_bool",
		"MaxSession_u32",
		"NoEnum_bool",
		"HubType_u32",
	}
}
