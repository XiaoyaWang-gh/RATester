package nathole

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestListAllLocalIPs_1(t *testing.T) {
	tests := []struct {
		name    string
		want    []net.IP
		wantErr bool
	}{
		{
			name:    "Test case 1",
			want:    []net.IP{net.ParseIP("127.0.0.1"), net.ParseIP("::1")},
			wantErr: false,
		},
		{
			name:    "Test case 2",
			want:    []net.IP{net.ParseIP("192.168.1.1"), net.ParseIP("fe80::1")},
			wantErr: false,
		},
		{
			name:    "Test case 3",
			want:    []net.IP{net.ParseIP("10.0.0.1"), net.ParseIP("fe80::1")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ListAllLocalIPs()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAllLocalIPs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllLocalIPs() = %v, want %v", got, tt.want)
			}
		})
	}
}
