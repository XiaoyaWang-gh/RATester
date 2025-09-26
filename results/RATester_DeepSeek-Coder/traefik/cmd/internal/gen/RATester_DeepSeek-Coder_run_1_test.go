package main

import (
	"fmt"
	"testing"
)

func TestRun_1(t *testing.T) {
	type args struct {
		dest string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test run with valid destination",
			args: args{
				dest: "testdata/dest",
			},
			wantErr: false,
		},
		{
			name: "Test run with invalid destination",
			args: args{
				dest: "invalid/dest",
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

			err := run(tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
