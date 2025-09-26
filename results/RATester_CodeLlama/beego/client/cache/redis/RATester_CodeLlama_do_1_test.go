package redis

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup your test
	rc := &Cache{}
	commandName := ""
	args := []interface{}{}
	want := interface{}(nil)
	wantErr := error(nil)
	got, err := rc.do(commandName, args...)
	if (err != nil) != (wantErr != nil) {
		t.Errorf("do() error = %v, wantErr %v", err, wantErr)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("do() = %v, want %v", got, want)
	}
}
