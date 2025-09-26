package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindataRead_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data := []byte("")
	name := ""
	want := []byte("")
	got, err := bindataRead(data, name)
	if err != nil {
		t.Errorf("bindataRead() error = %v, wantErr %v", err, false)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("bindataRead() = %v, want %v", got, want)
	}
}
