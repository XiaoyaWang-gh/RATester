package file

import (
	"fmt"
	"testing"
)

func TestGetHostId_1(t *testing.T) {
	db := &JsonDb{
		HostIncreaseId: 0,
	}

	tests := []struct {
		name string
		want int32
	}{
		{
			name: "TestGetHostId_1",
			want: 1,
		},
		{
			name: "TestGetHostId_2",
			want: 2,
		},
		{
			name: "TestGetHostId_3",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := db.GetHostId(); got != tt.want {
				t.Errorf("GetHostId() = %v, want %v", got, tt.want)
			}
		})
	}
}
