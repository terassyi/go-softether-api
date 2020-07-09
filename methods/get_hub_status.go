package methods

import (
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type GetHubStatus struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func NewGetHubStatus(name string) *GetHubStatus {
	return &GetHubStatus{
		Base:   pkg.NewBase("GetHubStatus"),
		Params: newGetHubStatusParams(name),
	}
}

func (m *GetHubStatus) Name() string {
	return m.Base.Name
}

func (m *GetHubStatus) GetId() int {
	return m.Id
}

func (m *GetHubStatus) SetId(id int) {
	m.Base.Id = id
}

func (m *GetHubStatus) Parameter() pkg.Params {
	return m.Params
}

func (m *GetHubStatus) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type GetHubStatusParams struct {
	Name string `json:"Name_str"`
}

func newGetHubStatusParams(name string) *GetHubStatusParams {
	return &GetHubStatusParams{
		Name: name,
	}
}

func (p *GetHubStatusParams) Tags() []string {
	return []string{
		"Name_str",
	}
}
