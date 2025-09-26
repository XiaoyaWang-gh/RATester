package logs

import (
	"fmt"
	"testing"
)

func TestSetGlobalFormatter_1(t *testing.T) {
	type args struct {
		fmtter string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				fmtter: "test",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				fmtter: "",
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

			if err := SetGlobalFormatter(tt.args.fmtter); (err != nil) != tt.wantErr {
				t.Errorf("SetGlobalFormatter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
