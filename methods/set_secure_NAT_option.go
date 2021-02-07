package methods

import (
	"encoding/json"

	"github.com/terassyi/go-softether-api/pkg"
)

type SetSecureNATOption struct {
	pkg.Base
	Params *SetSecureNATOptionParams `json:"params"`
}

func NewSetSecureNATOption(rpcHubName string) *SetSecureNATOption {
	return &SetSecureNATOption{
		Base:   pkg.NewBase("SetSecureNATOption"),
		Params: newSetSecureNATOptionParams(rpcHubName),
	}
}

func (m *SetSecureNATOption) Name() string {
	return m.Base.Name
}

func (m *SetSecureNATOption) GetId() int {
	return m.Id
}

func (m *SetSecureNATOption) SetId(id int) {
	m.Base.Id = id
}

func (m *SetSecureNATOption) Parameter() pkg.Params {
	return m.Params
}

func (m *SetSecureNATOption) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

type SetSecureNATOptionParams struct {
	RpcHubName            string `json:"RpcHubName_str"`
	MacAddressBin         string `json:"MacAddress_bin"`
	Ip                    string `json:"Ip_ip"`
	MaskIp                string `json:"Mask_ip"`
	UseNat                bool   `json:"UseNat_bool"`
	Mtu                   int    `json:"Mtu_u32"`
	NatTcpTimeout         int    `json:"NatTcpTimeout_u32"`
	NatUdpTimeout         int    `json:"NatUdpTimeout_u32"`
	UseDhcp               bool   `json:"UseDhcp_bool"`
	DhcpLeaseIPStart      string `json:"DhcpLeaseIPStart_ip"`
	DhcpLeaseIPEnd        string `json:"DhcpLeaseIPEnd_ip"`
	DhcpSubnetMask        string `json:"DhcpSubnetMask_ip"`
	DhcpExpireTimeSpan    int    `json:"DhcpExpireTimeSpan_u32"`
	DhcpGatewayAddress    string `json:"DhcpGatewayAddress_ip"`
	DhcpDnsServerAddress  string `json:"DhcpDnsServerAddress_ip"`
	DhcpDnsServerAddress2 string `json:"DhcpDnsServerAddress2_ip"`
	DhcpDomainName        string `json:"DhcpDomainName_str"`
	SaveLog               bool   `json:"SaveLog_bool"`
	ApplyDhcpPushRoutes   bool   `json:"ApplyDhcpPushRoutes_bool"`
	DhcpPushRoutes        string `json:"DhcpPushRoutes_str"`
}

func newSetSecureNATOptionParams(rpcHubName string) *SetSecureNATOptionParams {
	return &SetSecureNATOptionParams{
		RpcHubName: rpcHubName,
	}
}

func (p *SetSecureNATOptionParams) Tags() []string {
	return []string{
		"RpcHubName_str",
		"MacAddress_bin",
		"Ip_ip",
		"Mask_ip",
		"UseNat_bool",
		"Mtu_u32",
		"NatTcpTimeout_u32",
		"NatUdpTimeout_u32",
		"UseDhcp_bool",
		"DhcpLeaseIPStart_ip",
		"DhcpLeaseIPEnd_ip",
		"DhcpSubnetMask_ip",
		"DhcpExpireTimeSpan_u32",
		"DhcpGatewayAddress_ip",
		"DhcpDnsServerAddress_ip",
		"DhcpDnsServerAddress2_ip",
		"DhcpDomainName_str",
		"SaveLog_bool",
		"ApplyDhcpPushRoutes_bool",
		"DhcpPushRoutes_str",
	}
}
