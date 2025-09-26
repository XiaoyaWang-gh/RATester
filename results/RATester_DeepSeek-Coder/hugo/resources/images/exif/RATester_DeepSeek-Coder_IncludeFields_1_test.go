package exif

import (
	"fmt"
	"testing"
)

func TestIncludeFields_1(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				expression: "test",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				expression: "invalid regex",
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

			got := IncludeFields(tt.args.expression)
			d := &Decoder{}
			err := got(d)
			if (err != nil) != tt.wantErr {
				t.Errorf("IncludeFields() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
