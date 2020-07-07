package methods

import (
	"bytes"
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type EnumListener struct {
	pkg.Base
	Params pkg.Params `json:"params"`
}

func (g *EnumListener) Parameter() pkg.Params {
	return g.Params
}

func NewEnumListener() *EnumListener {
	return &EnumListener{
		Base:   pkg.NewBase("EnumListener"),
		Params: nil,
	}
}

func (g *EnumListener) Name() string {
	return g.Base.Name
}

func (g *EnumListener) GetId() int {
	return g.Id
}

func (g *EnumListener) SetId(id int) {
	g.Base.Id = id
}

func (g *EnumListener) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnumListenerParams struct{}

func (p *EnumListenerParams) Set(_key string, _val interface{}) error {
	return nil
}

func (p *EnumListenerParams) Get(_key string) (interface{}, error) {
	return nil, nil
}
