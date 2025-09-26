package allconfig

import (
	"errors"
	"fmt"
	"testing"
)

func TesttransientErr_1(t *testing.T) {
	tests := []struct {
		name    string
		c       *Configs
		wantErr bool
	}{
		{
			name: "Test case 1",
			c: &Configs{
				LanguageConfigSlice: []*Config{
					{
						C: &ConfigCompiled{
							transientErr: errors.New("test error"),
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test case 2",
			c: &Configs{
				LanguageConfigSlice: []*Config{
					{
						C: &ConfigCompiled{
							transientErr: nil,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			c: &Configs{
				LanguageConfigSlice: []*Config{},
			},
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

			if err := tt.c.transientErr(); (err != nil) != tt.wantErr {
				t.Errorf("transientErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
