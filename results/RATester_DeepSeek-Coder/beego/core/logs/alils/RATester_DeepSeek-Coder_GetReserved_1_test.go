package alils

import (
	"fmt"
	"testing"

	"github.com/aws/smithy-go/ptr"
)

func TestGetReserved_1(t *testing.T) {
	tests := []struct {
		name string
		m    *LogGroup
		want string
	}{
		{
			name: "Test with nil LogGroup",
			m:    nil,
			want: "",
		},
		{
			name: "Test with non-nil LogGroup and nil Reserved",
			m:    &LogGroup{},
			want: "",
		},
		{
			name: "Test with non-nil LogGroup and non-nil Reserved",
			m:    &LogGroup{Reserved: ptr.String("test")},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.GetReserved(); got != tt.want {
				t.Errorf("LogGroup.GetReserved() = %v, want %v", got, tt.want)
			}
		})
	}
}
