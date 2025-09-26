package allconfig

import (
	"errors"
	"fmt"
	"testing"
)

func TestTransientErr_1(t *testing.T) {
	tests := []struct {
		name    string
		c       *Configs
		wantErr bool
	}{
		{
			name: "Test with no transient error",
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
			name: "Test with transient error",
			c: &Configs{
				LanguageConfigSlice: []*Config{
					{
						C: &ConfigCompiled{
							transientErr: errors.New("transient error"),
						},
					},
				},
			},
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

			err := tt.c.transientErr()
			if (err != nil) != tt.wantErr {
				t.Errorf("transientErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
