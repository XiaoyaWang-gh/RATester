package mem

import (
	"fmt"
	"testing"
)

func TestAddTrafficIn_4(t *testing.T) {
	type args struct {
		name         string
		trafficBytes int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				name:         "proxy1",
				trafficBytes: 100,
			},
		},
		{
			name: "Test case 2",
			args: args{
				name:         "proxy2",
				trafficBytes: 200,
			},
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

			m := &serverMetrics{}
			m.AddTrafficIn(tt.args.name, "", tt.args.trafficBytes)

			// Add assertions here to verify the expected behavior
		})
	}
}
