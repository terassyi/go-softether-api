package methods

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/terassyi/go-softether-api/pkg"
)

type CreateLink struct {
	pkg.Base
	Params *CreateLinkParams `json:"params"`
}

func NewCreateLink(hubNameEx, accountName string) *CreateLink {
	return &CreateLink{
		Base:   pkg.NewBase("CreateLink"),
		Params: newCreateLinkParams(hubNameEx, accountName),
	}
}

func (m *CreateLink) Name() string {
	return m.Base.Name
}

func (m *CreateLink) GetId() int {
	return m.Id
}

func (m *CreateLink) SetId(id int) {
	m.Base.Id = id
}

func (m *CreateLink) Parameter() pkg.Params {
	return m.Params
}

func (m *CreateLink) Marshall() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type CreateLinkParams struct {
	HubNameEx                    string `json:"HubName_Ex_str"`
	CheckServerCert              bool   `json:"CheckServerCert_bool"`
	AccountName                  string `json:"AccountName_utf"`
	Hostname                     string `json:"Hostname_str"`
	Port                         int    `json:"Port_u32"`
	ProxyType                    int    `json:"ProxyType_u32"`
	HubName                      string `json:"HubName_str"`
	MaxConnection                int    `json:"MaxConnection_u32"`
	UseEncrypt                   bool   `json:"UseEncrypt_bool"`
	UseCompress                  bool   `json:"UseCompress_bool"`
	HalfConnection               bool   `json:"HalfConnection_bool"`
	AdditionalConnectionInterval int    `json:"AdditionalConnectionInterval_u32"`
	ConnectionDisconnectSpan     int    `json:"ConnectionDisconnectSpan_u32"`
	AuthType                     int    `json:"AuthType_u32"`
	Username                     string `json:"Username_str"`
	HashedPassword               string `json:"HashedPassword_bin"`
	PlainPassword                string `json:"PlainPassword_str"`
	ClientX                      string `json:"ClientX_bin"`
	ClientK                      string `json:"ClientK_bin"`
	CreateLinkPolicy
}

func newCreateLinkParams(hubNameEx, accountName string) *CreateLinkParams {
	return &CreateLinkParams{
		HubNameEx:   hubNameEx,
		AccountName: accountName,
	}
}

type CreateLinkPolicy struct {
	DHCPFilter              bool `json:"policy:DHCPFilter_bool"`
	DHCPNoServer            bool `json:"policy:DHCPNoServer_bool"`
	DHCPForce               bool `json:"policy:DHCPForce_bool"`
	CheckMac                bool `json:"SecPol_CheckMac_bool"`
	CheckIP                 bool `json:"SecPol_CheckIP_bool"`
	ArpDhcpOnly             bool `json:"policy:ArpDhcpOnly_bool"`
	PrivacyFilter           bool `json:"policy:PrivacyFilter_bool"`
	NoServer                bool `json:"policy:NoServer_bool"`
	NoBroadcastLimiter      bool `json:"policy:NoBroadcastLimiter_bool"`
	MaxMac                  int  `json:"policy:MaxMac_u32"`
	MaxIP                   int  `json:"policy:MaxIP_u32"`
	MaxUpload               int  `json:"policy:MaxUpload_u32"`
	MaxDownload             int  `json:"policy:MaxDownload_u32"`
	RSandRAFilter           bool `json:"policy:RSandRAFilter_bool"`
	RAFilter                bool `json:"SecPol_RAFilter_bool"`
	DHCPv6Filter            bool `json:"policy:DHCPv6Filter_bool"`
	DHCPv6NoServer          bool `json:"policy:DHCPv6NoServer_bool"`
	CheckIPv6               bool `json:"SecPol_CheckIPv6_bool"`
	NoServerV6              bool `json:"policy:NoServerV6_bool"`
	MaxIPv6                 int  `json:"policy:MaxIPv6_u32"`
	FilterIPv4              bool `json:"policy:FilterIPv4_bool"`
	FilterIPv6              bool `json:"policy:FilterIPv6_bool"`
	FilterNonIP             bool `json:"policy:FilterNonIP_bool"`
	NoIPv6DefaultRouterInRA bool `json:"policy:NoIPv6DefaultRouterInRA_bool"`
	VLanId                  int  `json:"policy:VLanId_u32"`
	Ver3                    bool `json:"policy:Ver3_bool"`
}

func (pol *CreateLinkPolicy) Tags() []string {
	tmp := CreateLinkPolicy{}
	t := reflect.TypeOf(tmp)
	var tags []string
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		tags = append(tags, tag)
	}
	return tags
}

func (p *CreateLinkParams) Tags() []string {
	tags := []string{
		"HubName_Ex_str",
		"CheckServerCert_bool",
		"AccountName_utf",
		"Hostname_str",
		"Port_u32",
		"ProxyType_u32",
		"HubName_str",
		"MaxConnection_u32",
		"UseEncrypt_bool",
		"UseCompress_bool",
		"HalfConnection_bool",
		"AdditionalConnectionInterval_u32",
		"ConnectionDisconnectSpan_u32",
		"AuthType_u32",
		"Username_str",
		"HashedPassword_bin",
		"PlainPassword_str",
		"ClientX_bin",
		"ClientK_bin",
	}
	tags = append(tags, p.CreateLinkPolicy.Tags()...)
	return tags
}
