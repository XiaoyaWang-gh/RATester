package context

import (
	"fmt"
	"testing"
)

func TestBindYAML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Input: &BeegoInput{
			RequestBody: []byte("name: beego\nage: 1\n"),
		},
	}
	var obj struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}
	err := ctx.BindYAML(&obj)
	if err != nil {
		t.Errorf("BindYAML failed. err: %v", err)
	}
	if obj.Name != "beego" {
		t.Errorf("BindYAML failed. obj.Name: %v", obj.Name)
	}
	if obj.Age != 1 {
		t.Errorf("BindYAML failed. obj.Age: %v", obj.Age)
	}
}
