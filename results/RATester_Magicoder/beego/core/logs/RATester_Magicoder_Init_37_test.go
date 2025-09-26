package logs

import (
	"fmt"
	"testing"
)

func TestInit_37(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		f       *multiFileLogWriter
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

			if err := tt.f.Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("multiFileLogWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
