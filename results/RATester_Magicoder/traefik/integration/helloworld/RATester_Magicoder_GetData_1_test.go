package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestGetData_1(t *testing.T) {
	tests := []struct {
		name string
		m    *StreamExampleReply
		want string
	}{
		{
			name: "Test case 1",
			m:    &StreamExampleReply{Data: "test data"},
			want: "test data",
		},
		{
			name: "Test case 2",
			m:    &StreamExampleReply{Data: ""},
			want: "",
		},
		{
			name: "Test case 3",
			m:    nil,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.GetData(); got != tt.want {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
