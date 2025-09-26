package commands

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestBuildErr_1(t *testing.T) {
	e := &hugoBuilderErrState{
		mu:       sync.Mutex{},
		paused:   false,
		builderr: nil,
		waserr:   false,
	}

	tests := []struct {
		name    string
		state   *hugoBuilderErrState
		wantErr bool
	}{
		{
			name:    "Test with no error",
			state:   e,
			wantErr: false,
		},
		{
			name: "Test with error",
			state: func() *hugoBuilderErrState {
				e.setBuildErr(errors.New("test error"))
				return e
			}(),
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

			err := tt.state.buildErr()
			if (err != nil) != tt.wantErr {
				t.Errorf("buildErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
