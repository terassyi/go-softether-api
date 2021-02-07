package methods

import (
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type AddLocalBridge struct {
	pkg.Base
	Params *AddLocalBridgeParams `json:"params"`
}

func NewAddLocalBridge(hubNameLB, deviceName string) *AddLocalBridge {
	return &AddLocalBridge{
		Base:   pkg.NewBase("AddLocalBridge"),
		Params: newAddLocalBridgeParams(hubNameLB, deviceName),
	}
}

func (m *AddLocalBridge) Name() string {
	return m.Base.Name
}

func (m *AddLocalBridge) GetId() int {
	return m.Id
}

func (m *AddLocalBridge) SetId(id int) {
	m.Base.Id = id
}

func (m *AddLocalBridge) Parameter() pkg.Params {
	return m.Params
}

func (m *AddLocalBridge) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type AddLocalBridgeParams struct {
	HubNameLB  string `json:"HubNameLB_str"`
	DeviceName string `json:"DeviceName_str"`
}

func newAddLocalBridgeParams(hubNameLB, deviceName string) *AddLocalBridgeParams {
	return &AddLocalBridgeParams{
		HubNameLB:  hubNameLB,
		DeviceName: deviceName,
	}
}

func (p *AddLocalBridgeParams) Tags() []string {
	return []string{
		"HubNameLB_str",
		"DeviceName_str",
	}
}
