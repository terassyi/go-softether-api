# SoftEther VPN Server JSON-RPC API in Golang
This is the SoftEther VPN Server JSON-RPC API wrapper in Go.
See a [document](https://github.com/SoftEtherVPN/SoftEtherVPN/tree/master/developer_tools/vpnserver-jsonrpc-clients/) for more information.

## support
supported API methods
- Test
- GetServerInfo
- GetServerStatus
- CreateListener
- EnumListener
- DeleteListener
- EnableListener
- SetServerPassword
- CreateHub
- SetHub
- EnumHub
- DeleteHub
- GetHubStatus
- SetHubOnline
- EnumConnection
- DisconnectConnection
- GetConnectionInfo
- CreateUser
- SetUser
- GetUser
- EnumUser

## Usage
To use this package, import this package like bellow.
```go
import softether "github.com/terassyi/go-softether-api"
```
First, you have to create a new Api instance.
```go
api := softether.New("localhost", 443, "default", "password")
```
Second, You can create api methods you want to call.
API methods are in methods package. These methods implement the `Method` interface.
For example, to call `Test` method.
```go
method := methods.NewTest()
```
And execute `api.Call()`. Then, result will return as `map[string](interface{})`.
```go
res, err := api.Call(method)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
```

## Future works
I'm making api methods as needed. Other methods will be supported.