package main

import (
	"fmt"
	"reflect"
	"testing"

	"fyne.io/fyne/v2"
)

func TestmakeMainTab_1(t *testing.T) {
	type args struct {
		serverPort string
		vKey       string
		connType   string
	}
	tests := []struct {
		name string
		args args
		want *fyne.Container
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := makeMainTab(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeMainTab() = %v, want %v", got, tt.want)
			}
		})
	}
}
