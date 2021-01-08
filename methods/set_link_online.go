package methods

import (
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type SetLinkOnline struct {
	pkg.Base
	Params *SetLinkOnlineParams `json:"params"`
}

func NewSetLinkOnline(hubname, accountname string) *SetLinkOnline {
	return &SetLinkOnline{
		Base:   pkg.NewBase("SetLinkOnline"),
		Params: newSetLinkOnlineParams(hubname, accountname),
	}
}

func (m *SetLinkOnline) Name() string {
	return m.Base.Name
}

func (m *SetLinkOnline) GetId() int {
	return m.Id
}

func (m *SetLinkOnline) SetId(id int) {
	m.Base.Id = id
}

func (m *SetLinkOnline) Parameter() pkg.Params {
	return m.Params
}

func (m *SetLinkOnline) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type SetLinkOnlineParams struct {
	HubName     string `json:"HubName_str"`
	AccountName string `json:"AccountName_utf"`
}

func newSetLinkOnlineParams(hubname, accountname string) *SetLinkOnlineParams {
	return &SetLinkOnlineParams{
		HubName:     hubname,
		AccountName: accountname,
	}
}

func (p *SetLinkOnlineParams) Tags() []string {
	return []string{
		"HubName_str",
		"AccountName_utf",
	}
}
