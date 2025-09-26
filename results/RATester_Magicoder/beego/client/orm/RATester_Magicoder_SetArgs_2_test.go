package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetArgs_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := rawSet{}
	args := []interface{}{1, "test", true}
	result := o.SetArgs(args...)
	if !reflect.DeepEqual(result.(*rawSet).args, args) {
		t.Errorf("Expected %v, got %v", args, result.(*rawSet).args)
	}
}
