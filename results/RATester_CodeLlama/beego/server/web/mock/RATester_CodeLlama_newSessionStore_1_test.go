package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSessionStore_1(t *testing.T) {
	type args struct {
		// TODO: Add test cases.
	}
	tests := []struct {
		name string
		args args
		want *SessionStore
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

			if got := newSessionStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSessionStore() = %v, want %v", got, tt.want)
			}
		})
	}
}
