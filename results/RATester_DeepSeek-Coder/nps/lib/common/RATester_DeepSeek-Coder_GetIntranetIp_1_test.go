package common

import (
	"fmt"
	"testing"
)

func TestGetIntranetIp_1(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		want    string
	}{
		{"Test case 1", false, "192.168.1.1"},
		{"Test case 2", true, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err, got := GetIntranetIp()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIntranetIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetIntranetIp() = %v, want %v", got, tt.want)
			}
		})
	}
}
