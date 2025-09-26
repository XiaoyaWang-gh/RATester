package web

import (
	"fmt"
	"html/template"
	"testing"
)

func TestRenderForm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string `form:"field1,text,id1,class1,true"`
		Field2 int    `form:"field2,number,id2,class2,false"`
	}

	testStruct := TestStruct{
		Field1: "test1",
		Field2: 123,
	}

	expected := template.HTML(`<label for="id1" class="class1">field1</label><input type="text" name="field1" id="id1" class="class1" value="test1" required></br><label for="id2" class="class2">field2</label><input type="number" name="field2" id="id2" class="class2" value="123">`)
	result := RenderForm(&testStruct)

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}
