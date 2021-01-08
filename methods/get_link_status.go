package methods

import (
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type GetLinkStatus struct {
	pkg.Base
	Params *GetLinkStatusParams `json:"params"`
}

func NewGetLinkStatus(hubname, accountname string) *GetLinkStatus {
	return &GetLinkStatus{
		Base:   pkg.NewBase("GetLinkStatus"),
		Params: newGetLinkStatusParams(hubname, accountname),
	}
}

func (m *GetLinkStatus) Name() string {
	return m.Base.Name
}

func (m *GetLinkStatus) GetId() int {
	return m.Id
}

func (m *GetLinkStatus) SetId(id int) {
	m.Base.Id = id
}

func (m *GetLinkStatus) Parameter() pkg.Params {
	return m.Params
}

func (m *GetLinkStatus) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type GetLinkStatusParams struct {
	HubNameEx   string `json:"HubName_Ex_str"`
	AccountName string `json:"AccountName_utf"`
}

func newGetLinkStatusParams(hubname, accountname string) *GetLinkStatusParams {
	return &GetLinkStatusParams{
		HubNameEx:   hubname,
		AccountName: accountname,
	}
}

func (p *GetLinkStatusParams) Tags() []string {
	return []string{
		"HubName_Ex_str",
		"AccountName_utf",
	}
}
