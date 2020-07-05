package methods

import "testing"

func TestParams_Get(t *testing.T) {
	p := Params{IntValue: 0}
	n, err := p.Get("IntValue")
	if err != nil {
		t.Fatal(err)
	}
	if n.(int) != 0 {
		t.Fatalf("diff: wanted:0 actual:%v", n)
	}
}

func TestParams_Set(t *testing.T) {
	p := Params{IntValue: 0}
	if err := p.Set("IntValue", 1); err != nil {
		t.Fatal(err)
	}
	if p.IntValue != 1 {
		t.Fatalf("diff: %d", p.IntValue)
	}
}

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
