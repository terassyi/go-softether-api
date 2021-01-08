package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type GetLink struct {
	pkg.Base
	Params *GetLinkParams `json:"params"`
}

func NewGetLink(hubNameEx, accountName string) *GetLink {
	return &GetLink{
		Base:   pkg.NewBase("GetLink"),
		Params: newGetLinkParams(hubNameEx, accountName),
	}
}

func (m *GetLink) Name() string {
	return m.Base.Name
}

func (m *GetLink) GetId() int {
	return m.Id
}

func (m *GetLink) SetId(id int) {
	m.Base.Id = id
}

func (m *GetLink) Parameter() pkg.Params {
	return m.Params
}

func (m *GetLink) Marshall() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type GetLinkParams struct {
	HubNameEx   string `json:"HubName_Ex_str"`
	AccountName string `json:"AccountName_utf"`
}

func newGetLinkParams(hubNameEx, accountName string) *GetLinkParams {
	return &GetLinkParams{HubNameEx: hubNameEx, AccountName: accountName}
}

func (p *GetLinkParams) Tags() []string {
	return []string{
		"HubName_Ex_str",
		"AccountName_utf",
	}
}
