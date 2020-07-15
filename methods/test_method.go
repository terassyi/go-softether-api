package methods

import (
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
)

type Test struct {
	pkg.Base
	Params *TestParams `json:"params"`
}

func NewTest() *Test {
	return &Test{
		Base:   pkg.NewBase("Test"),
		Params: &TestParams{IntValue: 0},
	}
}

func (m *Test) Name() string {
	return m.Base.Name
}

func (m *Test) GetId() int {
	return m.Id
}

func (m *Test) SetId(id int) {
	m.Base.Id = id
}

func (m *Test) Parameter() pkg.Params {
	return m.Params
}

func (m *Test) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type TestParams struct {
	IntValue int `json:"IntValue_u32"`
}

func (p *TestParams) Tags() []string {
	return []string{"IntValue_u32"}
}
