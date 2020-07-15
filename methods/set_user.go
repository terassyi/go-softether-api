package methods

import (
	"bytes"
	"encoding/json"
	"github.com/terassyi/go-softether-api/pkg"
	"reflect"
	"time"
)

type SetUSer struct {
	pkg.Base
	Params *SetUSerParams `json:"params"`
}

func (g *SetUSer) Parameter() pkg.Params {
	return g.Params
}

func NewSetUSer(hub, name, realname, group, note string, expire *time.Time, auth int) *SetUSer {
	return &SetUSer{
		Base:   pkg.NewBase("SetUSer"),
		Params: newSetUSerParmas(hub, name, realname, group, note, expire, auth),
	}
}

func (g *SetUSer) Name() string {
	return g.Base.Name
}

func (g *SetUSer) GetId() int {
	return g.Id
}

func (g *SetUSer) SetId(id int) {
	g.Base.Id = id
}

func (g *SetUSer) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type SetUSerParams struct {
	HubName        string     `json:"HubName_str"`
	Name           string     `json:"Name_str"`
	GroupName      string     `json:"GroupName_str"`
	Realname       string     `json:"Realname_utf"`
	Note           string     `json:"Note_utf"`
	ExpireTime     *time.Time `json:"ExpireTime_dt"`
	AuthType       int        `json:"AuthType_u32"`
	Auth_Passoword string     `json:"Auth_Password_str"`
	UserX          string     `json:"UserX_bin"`
	Serial         string     `json:"Serial_bin"`
	CommonName     string     `json:"CommonName_utf"`
	RadiusUsername string     `json:"RadiusUsername_utf"`
	NtUsername     string     `json:"NtUsername_str"`
	UsePolicy      bool       `josn:"UsePolicy_bool"`
	SetUSerPolicy
}

func newSetUSerParmas(hub, name, realname, group, note string, expire *time.Time, auth int) *SetUSerParams {
	return &SetUSerParams{
		HubName:    hub,
		Name:       name,
		Realname:   realname,
		GroupName:  group,
		Note:       note,
		ExpireTime: expire,
		AuthType:   auth,
	}
}

type SetUSerPolicy struct {
	Access                          bool `json:"policy:Access_bool"`
	DHCPFilter                      bool `json:"policy:DHCPFilter_bool"`
	DHCPNoServer                    bool `json:"policy:DHCPNoServer_bool"`
	DHCPForce                       bool `json:"policy:DHCPForce_bool"`
	NoBridge                        bool `json:"policy:NoBridge_bool"`
	NoRouting                       bool `json:"policy:NoRouting_bool"`
	CheckMac                        bool `json:"policy:CheckMac_bool"`
	CheckIp                         bool `json:"policy:CheckIp_bool"`
	ArpDhcpOnly                     bool `json:"policy:ArpDhcpOnly_bool"`
	PrivacyFilter                   bool `json:"policy:PrivacyFilter_bool"`
	NoServer                        bool `json:"policy:NoServer_bool"`
	NoBroadcastLimiter              bool `json:"policy:NoBroadcastLimiter_bool"`
	MonitorPort                     bool `json:"policy:MonitorPort_bool"`
	MaxConnection                   int  `json:"policy:MaxConnection_u32"`
	TimeOut                         int  `json:"policy:TimeOut_u32"`
	MaxMac                          int  `json:"policy:MaxMac_u32"`
	MaxIp                           int  `json:"policy:MaxIp_u32"`
	MaxUpload                       int  `json:"MaxUpload_u32"`
	MaxDownload                     int  `json:"MaxDownload_u32"`
	FixPassword                     bool `json:"policy:FixPassword_bool"`
	MultiLogins                     int  `json:"policy:MultiLogins_u32"`
	NoQos                           bool `json:"policy:NoQos_bool"`
	RSandRAFilter                   bool `json:"policy:RSandRAFilter_bool"`
	RAFilter                        bool `json:"policy:RAFilter_bool"`
	DHCPv6Filter                    bool `json:"policy:DHCPv6Filter_bool"`
	DHCPv6NoServer                  bool `json:"policy:DHCPv6NoServer_bool"`
	NoRoutingV6                     bool `json:"policy:NoRoutingV6_bool"`
	CheckIPv6                       bool `json:"policy:CheckIPv6_bool"`
	NoServerV6                      bool `json:"policy:NoServerV6_bool"`
	MaxIPv6                         int  `json:"policy:MaxIPv6_u32"`
	NoSavePassword                  bool `json:"policy:NoSavePassword_bool"`
	AutoDisconnect                  int  `json:"policy:AutoDisconnect_bool"`
	FilterIPv4                      bool `json:"policy:FilterIPv4_bool"`
	FilterIPv6                      bool `json:"policy:FilterIPv6_bool"`
	FilterNonIP                     bool `json:"policy:FilterNonIP_bool"`
	NoIPv6DefaultRouterInRA         bool `json:"policy:NoIPv6DefaultRouterInRA_bool"`
	NoIPv6DefaultRouterInRAWhenIPv6 bool `json:"policy:NoIPv6DefaultRouterInRAWhenIPv6_bool"`
	VLanId                          int  `json:"policy:VLanId_u32"`
	Ver3                            bool `json:"policy:policy:Ver3_bool"`
}

func (pol *SetUSerPolicy) Tags() []string {
	tmp := SetUSerPolicy{}
	t := reflect.TypeOf(tmp)
	var tags []string
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		tags = append(tags, tag)
	}
	return tags
}

func (p *SetUSerParams) Tags() []string {
	tags := []string{
		"HubName_str",
		"Name_str",
		"Realname_utf",
		"Note_utf",
		"ExpireTime_dt",
		"AuthType_u32",
		"Auth_Password_str",
		"UserX_bin",
		"Serial_bin",
		"CommonName_utf",
		"RadiusUsername_utf",
		"NtUsername_utf",
		"UsePolicy_bool",
	}
	tags = append(tags, p.SetUSerPolicy.Tags()...)
	return tags
}
