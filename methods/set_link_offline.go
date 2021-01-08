package methods

import (
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type SetLinkOffline struct {
	pkg.Base
	Params *SetLinkOfflineParams `json:"params"`
}

func NewSetLinkOffline(hubname, accountname string) *SetLinkOffline {
	return &SetLinkOffline{
		Base:   pkg.NewBase("SetLinkOffline"),
		Params: newSetLinkOfflineParams(hubname, accountname),
	}
}

func (m *SetLinkOffline) Name() string {
	return m.Base.Name
}

func (m *SetLinkOffline) GetId() int {
	return m.Id
}

func (m *SetLinkOffline) SetId(id int) {
	m.Base.Id = id
}

func (m *SetLinkOffline) Parameter() pkg.Params {
	return m.Params
}

func (m *SetLinkOffline) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type SetLinkOfflineParams struct {
	HubName     string `json:"HubName_str"`
	AccountName string `json:"AccountName_utf"`
}

func newSetLinkOfflineParams(hubname, accountname string) *SetLinkOfflineParams {
	return &SetLinkOfflineParams{
		HubName:     hubname,
		AccountName: accountname,
	}
}

func (p *SetLinkOfflineParams) Tags() []string {
	return []string{
		"HubName_str",
		"AccountName_utf",
	}
}
