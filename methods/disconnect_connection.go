package methods

import (
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type DisconnectConnection struct {
	pkg.Base
	Params *DisconnectConnectionParams `json:"params"`
}

func NewDisconnectConnection(name string) *DisconnectConnection {
	return &DisconnectConnection{
		Base:   pkg.NewBase("DisconnectConnection"),
		Params: newDisconnectConnectionParams(name),
	}
}

func (m *DisconnectConnection) Name() string {
	return m.Base.Name
}

func (m *DisconnectConnection) GetId() int {
	return m.Id
}

func (m *DisconnectConnection) SetId(id int) {
	m.Base.Id = id
}

func (m *DisconnectConnection) Parameter() pkg.Params {
	return m.Params
}

func (m *DisconnectConnection) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type DisconnectConnectionParams struct {
	Name string `json:"Name_str"`
}

func newDisconnectConnectionParams(name string) *DisconnectConnectionParams {
	return &DisconnectConnectionParams{
		Name: name,
	}
}

func (p *DisconnectConnectionParams) Tags() []string {
	return []string{
		"Name_str",
	}
}
