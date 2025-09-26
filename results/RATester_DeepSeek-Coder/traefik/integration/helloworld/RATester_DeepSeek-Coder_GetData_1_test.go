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
			name: "Test with non-nil StreamExampleReply",
			m:    &StreamExampleReply{Data: "testData"},
			want: "testData",
		},
		{
			name: "Test with nil StreamExampleReply",
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
				t.Errorf("StreamExampleReply.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
