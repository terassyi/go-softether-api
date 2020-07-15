package methods

import (
	"fmt"
	"testing"
)

func TestCreateUserParams_Tags(t *testing.T) {
	m := NewCreateUser("test", "test_user", "hoge", "test", nil, 1)
	tags := m.Params.Tags()
	fmt.Println(tags)
}
