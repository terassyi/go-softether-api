package pkg

import (
	"fmt"
	"testing"
)

type testParams struct {
	Field1    string `json:"Field1_str"`
	Field2    int    `json:"Field2_u32"`
	Field3    bool   `json:"Field3_bool"`
	Field_a   bool   `json:"Field_a_bool"`
	FieldList []int  `json:"FieldList"`
}

func (p *testParams) Empty() *testParams {
	return &testParams{}
}

func (p *testParams) Tags() []string {
	return []string{
		"Field1_str",
		"Field2_u32",
		"Field3_bool",
		"Field_a_bool",
		"FieldList",
	}
}

func (p *testParams) Set(key string, val interface{}) error {
	//n := &testParams{}
	//m, err := StructToMap(p)
	//if err != nil {
	//	return err
	//}
	//m[key] = val
	//if err = MapToStruct(m, n); err != nil {
	//	return err
	//}
	//*p = *n
	switch key {
	case "Field1":
		p.Field1 = val.(string)
	case "Field2":
		p.Field2 = val.(int)
	case "Field3":
		p.Field3 = val.(bool)
	case "Field_a":
		p.Field_a = val.(bool)
	case "FieldList":
		p.FieldList = val.([]int)
	default:
		return fmt.Errorf("no such parameter")

	}
	return nil
}

func (p *testParams) Get(key string) (interface{}, error) {
	data, err := StructToMap(p)
	if err != nil {
		return nil, err
	}
	if data[key] == nil {
		return nil, fmt.Errorf("no such parameter")
	}
	return data[key], nil
}

func TestStructToMap(t *testing.T) {
	p := &testParams{
		Field1: "field1",
		Field2: 1,
		Field3: false,
	}
	wanted := make(map[string]interface{})
	wanted["Field1"] = p.Field1
	wanted["Field2"] = p.Field2
	wanted["Field3"] = p.Field3
	m, err := StructToMap(p)
	if err != nil {
		t.Fatal(err)
	}
	if m["Field1"] != wanted["Field1"] {
		t.Fatalf("actual: %v", m)
	}
}

func TestStructToMapNoField(t *testing.T) {
	p := &testParams{
		Field1: "field1",
		Field2: 1,
		Field3: false,
	}
	wanted := make(map[string]interface{})
	wanted["Field1"] = p.Field1
	wanted["Field2"] = p.Field2
	wanted["Field3"] = p.Field3
	wanted["Field4"] = "dont exist"
	fmt.Println(wanted)
	m, err := StructToMap(p)
	if err != nil {
		t.Fatal(err)
	}
	if m["Field1"] != wanted["Field1"] {
		t.Fatalf("actual: %v", m)
	}
	if m["Field4"] != nil {
		t.Fatalf("found invalid field")
	}
}

func TestSplitType(t *testing.T) {
	p := &testParams{
		Field1:    "field1",
		Field2:    1,
		Field3:    false,
		Field_a:   true,
		FieldList: []int{0, 1},
	}
	d, err := StructToMap(p)
	if err != nil {
		t.Fatal(err)
	}
	if d["Field_a"] != true {
		t.Fatalf("actual: %v", d)
	}
	fmt.Println(d)
	tmp := d["FieldList"].([]interface{})
	v := tmp[1]
	wanted := 1
	if int(v.(float64)) != 1 {
		t.Fatalf("failed - wanted: %v - actual: %v", wanted, v)
	}
}

func TestMapToStruct(t *testing.T) {
	p := &testParams{
		Field1:    "field1",
		Field2:    1,
		Field3:    false,
		Field_a:   true,
		FieldList: []int{0, 1},
	}
	res, err := p.Get("Field2")
	if err != nil {
		t.Fatal(err)
	}
	resInt := int(res.(float64))
	wanted := 1
	if resInt != wanted {
		t.Fatalf("actual: %v", res)
	}
}

func TestMapToStructModify(t *testing.T) {
	p := &testParams{
		Field1:    "field1",
		Field2:    1,
		Field3:    false,
		Field_a:   true,
		FieldList: []int{0, 1},
	}
	if err := p.Set("FieldList", []int{1, 2}); err != nil {
		t.Fatal(err)
	}
	res, err := p.Get("FieldList")
	if err != nil {
		t.Fatal(err)
	}
	resString := res.([]int)
	wanted := []int{1, 2}
	if resString[0] != wanted[0] {
		t.Fatalf("actual: %v", resString)
	}
}
