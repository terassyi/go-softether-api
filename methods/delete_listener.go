package methods

import (
	"bytes"
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type DeleteListener struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func (g *DeleteListener) Parameter() pkg.Params {
	return g.Params
}

func NewDeleteListener(port int) *DeleteListener {
	return &DeleteListener{
		Base:   pkg.NewBase("DeleteListener"),
		Params: newDeleteListenerParams(port),
	}
}

func (g *DeleteListener) Name() string {
	return g.Base.Name
}

func (g *DeleteListener) GetId() int {
	return g.Id
}

func (g *DeleteListener) SetId(id int) {
	g.Base.Id = id
}

func (g *DeleteListener) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type DeleteListenerParams struct {
	Port int `json:"Port_u32"`
}

func newDeleteListenerParams(port int) *DeleteListenerParams {
	return &DeleteListenerParams{Port: port}
}

func (p *DeleteListenerParams) Tags() []string {
	return []string{"Port_u32"}
}
