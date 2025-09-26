package filenotify

import (
	"fmt"
	"testing"
)

func Testremove_2(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		w       *filePoller
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.w.remove(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("filePoller.remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
