package methods

import (
	"bytes"
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type EnumConnection struct {
	pkg.Base
	Params *EnumConnectionParams `json:"params"`
}

func (g *EnumConnection) Parameter() pkg.Params {
	return g.Params
}

func NewEnumConnection() *EnumConnection {
	return &EnumConnection{
		Base:   pkg.NewBase("EnumConnection"),
		Params: &EnumConnectionParams{},
	}
}

func (g *EnumConnection) Name() string {
	return g.Base.Name
}

func (g *EnumConnection) GetId() int {
	return g.Id
}

func (g *EnumConnection) SetId(id int) {
	g.Base.Id = id
}

func (g *EnumConnection) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnumConnectionParams struct{}

func (p *EnumConnectionParams) Tags() []string {
	return []string{}
}
