package methods

import (
	"bytes"
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type EnumHub struct {
	pkg.Base
	Params *EnumHubParams `json:"params"`
}

func (g *EnumHub) Parameter() pkg.Params {
	return g.Params
}

func NewEnumHub() *EnumHub {
	return &EnumHub{
		Base:   pkg.NewBase("EnumHub"),
		Params: &EnumHubParams{},
	}
}

func (g *EnumHub) Name() string {
	return g.Base.Name
}

func (g *EnumHub) GetId() int {
	return g.Id
}

func (g *EnumHub) SetId(id int) {
	g.Base.Id = id
}

func (g *EnumHub) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnumHubParams struct{}

func (p *EnumHubParams) Tags() []string {
	return []string{}
}
