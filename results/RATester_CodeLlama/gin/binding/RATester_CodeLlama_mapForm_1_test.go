package binding

import (
	"fmt"
	"testing"
)

func TestMapForm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type User struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}
	user := User{}
	form := map[string][]string{
		"name": {"test"},
		"age":  {"18"},
	}
	err := mapForm(&user, form)
	if err != nil {
		t.Errorf("mapForm error: %v", err)
	}
	if user.Name != "test" {
		t.Errorf("user.Name = %v, want %v", user.Name, "test")
	}
	if user.Age != 18 {
		t.Errorf("user.Age = %v, want %v", user.Age, 18)
	}
}
