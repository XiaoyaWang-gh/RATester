package migration

import (
	"fmt"
	"testing"
)

func TestGetSQL_1(t *testing.T) {
	type args struct {
		m *Migration
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := tt.args.m.GetSQL(); got != tt.want {
				t.Errorf("GetSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
