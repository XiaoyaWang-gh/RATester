package filenotify

import (
	"fmt"
	"testing"
)

func Testrecord_1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		r       *recording
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

			if err := tt.r.record(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("recording.record() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
