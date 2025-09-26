package logs

import (
	"fmt"
	"testing"
)

func TestSetLogger_2(t *testing.T) {
	type args struct {
		adapter string
		config  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				adapter: "file",
				config:  []string{"logs/test.log"},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				adapter: "unknown",
				config:  []string{"logs/test.log"},
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

			err := SetLogger(tt.args.adapter, tt.args.config...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetLogger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
