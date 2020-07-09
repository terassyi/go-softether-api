package methods

import (
	"bytes"
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type GetServerStatus struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func (g *GetServerStatus) Parameter() pkg.Params {
	return g.Params
}

func NewGetServerStatus() *GetServerStatus {
	return &GetServerStatus{
		Base:   pkg.NewBase("GetServerStatus"),
		Params: nil,
	}
}

func (g *GetServerStatus) Name() string {
	return g.Base.Name
}

func (g *GetServerStatus) GetId() int {
	return g.Id
}

func (g *GetServerStatus) SetId(id int) {
	g.Base.Id = id
}

func (g *GetServerStatus) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type GetServerStatusParams struct{}

func (p *GetServerStatusParams) Tags() []string {
	return []string{}
}
