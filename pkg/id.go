package pkg

// Id interface is used for request id
type Id interface {
	Incl()
	Describe() int
}

// id struct is the implement of Id interface
type id int

// This method increment an id
func (i *id) Incl() {
	*i += 1
}

func (i *id) Describe() int {
	return int(*i)
}
