package methods

import (
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type GetConnectionInfo struct {
	pkg.Base
	Params *GetConnectionInfoParams `json:"params"`
}

func NewGetConnectionInfo(name string) *GetConnectionInfo {
	return &GetConnectionInfo{
		Base:   pkg.NewBase("GetConnectionInfo"),
		Params: newGetConnectionInfoParams(name),
	}
}

func (m *GetConnectionInfo) Name() string {
	return m.Base.Name
}

func (m *GetConnectionInfo) GetId() int {
	return m.Id
}

func (m *GetConnectionInfo) SetId(id int) {
	m.Base.Id = id
}

func (m *GetConnectionInfo) Parameter() pkg.Params {
	return m.Params
}

func (m *GetConnectionInfo) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type GetConnectionInfoParams struct {
	Name string `json:"Name_str"`
}

func newGetConnectionInfoParams(name string) *GetConnectionInfoParams {
	return &GetConnectionInfoParams{
		Name: name,
	}
}

func (p *GetConnectionInfoParams) Tags() []string {
	return []string{
		"Name_str",
	}
}
