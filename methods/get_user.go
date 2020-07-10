package methods

import (
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type GetUser struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func (m *GetUser) Parameter() pkg.Params {
	return m.Params
}

func NewGetUser(hub, name string) *GetUser {
	return &GetUser{
		Base:   pkg.NewBase("GetUser"),
		Params: newGetUserParams(hub, name),
	}
}

func (m *GetUser) Name() string {
	return m.Base.Name
}

func (m *GetUser) GetId() int {
	return m.Id
}

func (m *GetUser) SetId(id int) {
	m.Base.Id = id
}

func (m *GetUser) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type GetUserParams struct {
	HubName string `json:"HubName_str"`
	Name    string `json:"Name_str"`
}

func newGetUserParams(hub, name string) *GetUserParams {
	return &GetUserParams{
		HubName: hub,
		Name:    name,
	}
}

func (p *GetUserParams) Tags() []string {
	return []string{
		"HubName_str",
		"Name_str",
	}
}
