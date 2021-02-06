package methods

import (
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type DeleteLocalBridge struct {
	pkg.Base
	Params *DeleteLocalBridgeParams `json:"params"`
}

func NewDeleteLocalBridge(hubNameLB, deviceName string) *DeleteLocalBridge {
	return &DeleteLocalBridge{
		Base:   pkg.NewBase("DeleteLocalBridge"),
		Params: newDeleteLocalBridgeParams(hubNameLB, deviceName),
	}
}

func (m *DeleteLocalBridge) Name() string {
	return m.Base.Name
}

func (m *DeleteLocalBridge) GetId() int {
	return m.Id
}

func (m *DeleteLocalBridge) SetId(id int) {
	m.Base.Id = id
}

func (m *DeleteLocalBridge) Parameter() pkg.Params {
	return m.Params
}

func (m *DeleteLocalBridge) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type DeleteLocalBridgeParams struct {
	HubNameLB  string `json:"HubNameLB_str"`
	DeviceName string `json:"DeviceName_str"`
}

func newDeleteLocalBridgeParams(hubNameLB, deviceName string) *DeleteLocalBridgeParams {
	return &DeleteLocalBridgeParams{
		HubNameLB:  hubNameLB,
		DeviceName: deviceName,
	}
}

func (p *DeleteLocalBridgeParams) Tags() []string {
	return []string{
		"HubNameLB_str",
		"DeviceName_str",
	}
}
