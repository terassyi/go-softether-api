package methods

import "testing"

func TestTest_Name(t *testing.T) {
	p := NewTest()
	if p.Name() != "Test" {
		t.Fatalf("actual: %v", p.Name())
	}
}

func TestTest_Marshall(t *testing.T) {
	r := NewTest()
	b, err := r.Marshall()
	if err != nil {
		t.Fatal(err)
	}
	wanted := "{\"jsonrpc\":\"2.0\",\"id\":0,\"method\":\"Test\",\"params\":{\"IntValue_u32\":0}}"
	if string(b) != wanted {
		t.Fatalf("actual: %v", string(b))
	}
}
