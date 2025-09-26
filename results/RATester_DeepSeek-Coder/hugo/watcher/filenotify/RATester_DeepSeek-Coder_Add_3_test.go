package filenotify

import (
	"fmt"
	"testing"
)

func TestAdd_3(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		poller  *filePoller
		wantErr bool
	}{
		{
			name: "success",
			poller: &filePoller{
				watches: make(map[string]struct{}),
			},
			wantErr: false,
		},
		{
			name: "poller closed",
			poller: &filePoller{
				watches: make(map[string]struct{}),
				closed:  true,
			},
			wantErr: true,
		},
		{
			name: "watch exists",
			poller: &filePoller{
				watches: map[string]struct{}{
					"test": {},
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

			err := tt.poller.Add("test")
			if (err != nil) != tt.wantErr {
				t.Errorf("filePoller.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
