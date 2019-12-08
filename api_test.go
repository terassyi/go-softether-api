package softetherapi

import (
	"fmt"
	"testing"
)

func TestGeneratePass(t *testing.T) {
	result := generatePassword()
	fmt.Println(result)
}

func TestApiMethodToMap(t *testing.T) {
	c := CreateUSerParams{
		HubName:           "VPN",
		Name_str:          "test",
		Realname_str:      "test",
		Note_utf:          "hogehoge",
		ExpireTime_dt:     "2019:12:02T12:00",
		AuthType_u32:      1,
		Auth_Password_str: "password",
	}
	i := apiId(0)
	i.call()
	req := apiMethod{
		baseField{
			jsonrpc: Jsonrpc,
			id:      i,
			method:  "CreateUser",
		},
		c,
	}
	a, err := req.ToMap()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a)
}

func TestToJSON(t *testing.T) {
	c := CreateUSerParams{
		HubName:           "VPN",
		Name_str:          "test",
		Realname_str:      "test",
		Note_utf:          "hogehoge",
		ExpireTime_dt:     "2019:12:02T12:00",
		AuthType_u32:      1,
		Auth_Password_str: "password",
	}
	i := apiId(0)
	i.call()
	req := apiMethod{
		baseField{
			jsonrpc: Jsonrpc,
			id:      i,
			method:  "CreateUser",
		},
		c,
	}
	jsonReq, err := req.ToJSON()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonReq))

}
