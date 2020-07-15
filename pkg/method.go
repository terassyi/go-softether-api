package pkg

// Method interface is type for API request
type Method interface {
	Name() string
	GetId() int
	SetId(id int)
	Parameter() Params
	Marshall() ([]byte, error)
}

type Params interface {
	Tags() []string
}

const (
	ver = "2.0"
)

type Base struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Name    string `json:"method"`
}

func NewBase(name string) Base {
	return Base{
		Jsonrpc: ver,
		Id:      0,
		Name:    name,
	}
}
