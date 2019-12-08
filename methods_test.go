package softetherapi

import (
	"fmt"
	"testing"
)

func TestToMap(t *testing.T) {
	c := CreateUserParams{
		HubName:           "VPN",
		Name_str:          "test",
		Realname_str:      "test",
		Note_utf:          "hogehoge",
		ExpireTime_dt:     "2019:12:02T12:00",
		AuthType_u32:      1,
		Auth_Password_str: "password",
	}
	b := c.ToMap()
	fmt.Println(b)
}
