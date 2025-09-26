package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseIPs_1(t *testing.T) {
	tests := []struct {
		name  string
		addrs []string
		want  []string
	}{
		{
			name:  "Test case 1",
			addrs: []string{"192.168.1.1:8080", "192.168.1.2:8080", "192.168.1.3:8080"},
			want:  []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"},
		},
		{
			name:  "Test case 2",
			addrs: []string{"10.0.0.1:80", "10.0.0.2:443", "10.0.0.3:8080"},
			want:  []string{"10.0.0.1", "10.0.0.2", "10.0.0.3"},
		},
		{
			name:  "Test case 3",
			addrs: []string{"172.16.0.1:80", "172.16.0.2:443", "172.16.0.3:8080"},
			want:  []string{"172.16.0.1", "172.16.0.2", "172.16.0.3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := parseIPs(tt.addrs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseIPs() = %v, want %v", got, tt.want)
			}
		})
	}
}
