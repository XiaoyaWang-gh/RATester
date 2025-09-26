package common

import (
	"fmt"
	"testing"
)

func TestGetBufPoolCopy_1(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{
			name: "Test case 1",
			want: make([]byte, PoolSizeCopy),
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := GetBufPoolCopy()
			if len(got) != len(tt.want) {
				t.Errorf("GetBufPoolCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}
