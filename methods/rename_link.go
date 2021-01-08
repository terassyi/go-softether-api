package methods

import (
	"bytes"
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type RenameLink struct {
	pkg.Base
	Params *RenameLinkParams `json:"params"`
}

func NewRenameLink(hubname, oldAccountname, newAccountname string) *RenameLink {
	return &RenameLink{
		Base:   pkg.NewBase("RenameLink"),
		Params: newRenameLinkParams(hubname, oldAccountname, newAccountname),
	}
}
func (m *RenameLink) Name() string {
	return m.Base.Name
}

func (m *RenameLink) GetId() int {
	return m.Id
}

func (m *RenameLink) SetId(id int) {
	m.Base.Id = id
}

func (m *RenameLink) Parameter() pkg.Params {
	return m.Params
}

func (m *RenameLink) Marshall() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type RenameLinkParams struct {
	HubName        string `json:"HubName_str"`
	OldAccountName string `json:"OldAccountName_utf"`
	NewAccountName string `json:"NewAccountName_utf"`
}

func newRenameLinkParams(hubname, oldAccountname, newAccountname string) *RenameLinkParams {
	return &RenameLinkParams{HubName: hubname, OldAccountName: oldAccountname, NewAccountName: newAccountname}
}

func (p *RenameLinkParams) Tags() []string {
	return []string{
		"HubName_str",
		"OldAccountName_utf",
		"NewAccountName_utf",
	}
}
