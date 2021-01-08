package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type DeleteLink struct {
	pkg.Base
	Params *DeleteLinkParams `json:"params"`
}

func NewDeleteLink(hubname, accountname string) *DeleteLink {
	return &DeleteLink{
		Base:   pkg.NewBase("DeleteLink"),
		Params: newDeleteLinkParams(hubname, accountname),
	}
}
func (m *DeleteLink) Name() string {
	return m.Base.Name
}

func (m *DeleteLink) GetId() int {
	return m.Id
}

func (m *DeleteLink) SetId(id int) {
	m.Base.Id = id
}

func (m *DeleteLink) Parameter() pkg.Params {
	return m.Params
}

func (m *DeleteLink) Marshall() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type DeleteLinkParams struct {
	HubName     string `json:"HubName_str"`
	AccountName string `json:"AccountName_utf"`
}

func newDeleteLinkParams(hubname, accountname string) *DeleteLinkParams {
	return &DeleteLinkParams{HubName: hubname, AccountName: accountname}
}

func (p *DeleteLinkParams) Tags() []string {
	return []string{
		"HubName_str",
		"AccountName_utf",
	}
}
