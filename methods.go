package softetherapi

import "fmt"

type CreateUserParams struct {
	HubName_str       string
	Name_str          string
	Realname_utf      string
	Note_utf          string
	ExpireTime_dt     string
	AuthType_u32      int
	Auth_Password_str string
}

func (p CreateUserParams) Show() {
	fmt.Println("Method: CreateUser")
	fmt.Println("HubName_str: ", p.HubName_str)
	fmt.Println("Name_str: ", p.Name_str)
	fmt.Println("Realname_str: ", p.Realname_utf)
	fmt.Println("Note_utf: ", p.Note_utf)
	fmt.Println("ExpireTime_dt: ", p.ExpireTime_dt)
	fmt.Println("AuthType_u32: ", p.AuthType_u32)
	fmt.Println("Auth_Password_str: ", p.Auth_Password_str)
}

func (p CreateUserParams) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["HubName_str"] = p.HubName_str
	m["Name_str"] = p.Name_str
	m["Realname_utf"] = p.Realname_utf
	m["Note_utf"] = p.Note_utf
	m["ExpireTime_dt"] = p.ExpireTime_dt
	m["AuthType_u32"] = p.AuthType_u32
	m["Auth_Password_str"] = p.Auth_Password_str
	return m
}

func NewCreateUserParams(hub, name, password, note, dead string) CreateUserParams {
	return CreateUserParams{
		HubName_str:       hub,
		Name_str:          name,
		Realname_utf:      name,
		Note_utf:          note,
		ExpireTime_dt:     dead,
		AuthType_u32:      AuthType_u32,
		Auth_Password_str: password,
	}
}
