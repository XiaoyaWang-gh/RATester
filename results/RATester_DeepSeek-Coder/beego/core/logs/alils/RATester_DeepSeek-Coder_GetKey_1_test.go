package alils

import (
	"fmt"
	"testing"

	"github.com/aws/smithy-go/ptr"
)

func TestGetKey_1(t *testing.T) {
	tests := []struct {
		name string
		m    *LogContent
		want string
	}{
		{
			name: "Test with nil LogContent",
			m:    nil,
			want: "",
		},
		{
			name: "Test with non-nil LogContent",
			m:    &LogContent{Key: ptr.String("testKey")},
			want: "testKey",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.GetKey(); got != tt.want {
				t.Errorf("LogContent.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
