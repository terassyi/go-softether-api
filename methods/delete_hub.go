package methods

import (
	"bytes"
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type DeleteHub struct {
	pkg.Base
	Params *DeleteHubParams `json:"params"`
}

func NewDeleteHub(name string) *DeleteHub {
	return &DeleteHub{
		Base:   pkg.NewBase("DeleteHub"),
		Params: newDeleteHubParams(name),
	}
}
func (m *DeleteHub) Name() string {
	return m.Base.Name
}

func (m *DeleteHub) GetId() int {
	return m.Id
}

func (m *DeleteHub) SetId(id int) {
	m.Base.Id = id
}

func (m *DeleteHub) Parameter() pkg.Params {
	return m.Params
}

func (m *DeleteHub) Marshall() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type DeleteHubParams struct {
	HubName string `json:"HubName_str"`
}

func newDeleteHubParams(hub string) *DeleteHubParams {
	return &DeleteHubParams{HubName: hub}
}

func (p *DeleteHubParams) Tags() []string {
	return []string{
		"HubName_str",
	}
}
