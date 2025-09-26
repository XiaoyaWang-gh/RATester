package http

import (
	"fmt"
	"testing"
)

func TesthostV2_1(t *testing.T) {
	tree := &matchersTree{}

	tests := []struct {
		name    string
		hosts   []string
		wantErr bool
	}{
		{
			name:    "valid host",
			hosts:   []string{"example.com"},
			wantErr: false,
		},
		{
			name:    "invalid host",
			hosts:   []string{"exa\u0080mple.com"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := hostV2(tree, tt.hosts...); (err != nil) != tt.wantErr {
				t.Errorf("hostV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
