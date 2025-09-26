package fiber

import (
	"fmt"
	"testing"
)

func TestListen_1(t *testing.T) {
	t.Parallel()

	app := &App{}

	testCases := []struct {
		name    string
		addr    string
		config  ListenConfig
		wantErr bool
	}{
		{
			name:   "success",
			addr:   ":3000",
			config: ListenConfig{},
		},
		{
			name:    "failure",
			addr:    ":3000",
			config:  ListenConfig{CertFile: "non-existent-file"},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := app.Listen(tc.addr, tc.config)
			if (err != nil) != tc.wantErr {
				t.Errorf("Listen() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
