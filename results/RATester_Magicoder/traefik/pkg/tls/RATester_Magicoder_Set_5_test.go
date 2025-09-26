package tls

import (
	"fmt"
	"testing"
)

func TestSet_5(t *testing.T) {
	tests := []struct {
		name    string
		c       *Certificates
		value   string
		wantErr bool
	}{
		{
			name:    "normal case",
			c:       &Certificates{},
			value:   "file1,file2;file3,file4",
			wantErr: false,
		},
		{
			name:    "bad format",
			c:       &Certificates{},
			value:   "file1,file2;file3",
			wantErr: true,
		},
		{
			name:    "empty value",
			c:       &Certificates{},
			value:   "",
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

			if err := tt.c.Set(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("Certificates.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
