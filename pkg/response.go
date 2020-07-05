package pkg

type Response interface {
	Id() int
	Result() (map[string]interface{}, error)
}
