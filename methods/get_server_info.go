package methods

import (
	"bytes"
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type GetServerInfo struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func (g *GetServerInfo) Parameter() pkg.Params {
	return g.Params
}

func NewGetServerInfo() *GetServerInfo {
	return &GetServerInfo{
		Base:   pkg.NewBase("GetServerInfo"),
		Params: nil,
	}
}

func (g *GetServerInfo) Name() string {
	return g.Base.Name
}

func (g *GetServerInfo) GetId() int {
	return g.Id
}

func (g *GetServerInfo) SetId(id int) {
	g.Base.Id = id
}

func (g *GetServerInfo) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type GetServerInfoParams struct{}

func (p *GetServerInfoParams) Set() error {
	return nil
}

func (p *GetServerInfoParams) Get() (interface{}, error) {
	return nil, nil
}
