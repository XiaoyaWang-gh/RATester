package grace

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

func TestRegisterSignalHook_1(t *testing.T) {
	srv := &Server{
		SignalHooks: make(map[int]map[os.Signal][]func()),
	}

	testFunc := func() {}

	tests := []struct {
		name    string
		ppFlag  int
		sig     os.Signal
		f       func()
		wantErr bool
	}{
		{
			name:    "valid pre signal",
			ppFlag:  PreSignal,
			sig:     syscall.SIGINT,
			f:       testFunc,
			wantErr: false,
		},
		{
			name:    "valid post signal",
			ppFlag:  PostSignal,
			sig:     syscall.SIGINT,
			f:       testFunc,
			wantErr: false,
		},
		{
			name:    "invalid ppFlag",
			ppFlag:  2,
			sig:     syscall.SIGINT,
			f:       testFunc,
			wantErr: true,
		},
		{
			name:    "unsupported signal",
			ppFlag:  PreSignal,
			sig:     syscall.Signal(0x1),
			f:       testFunc,
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

			err := srv.RegisterSignalHook(tt.ppFlag, tt.sig, tt.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterSignalHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
