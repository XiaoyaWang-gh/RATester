package common

import (
	"fmt"
	"testing"
)

func TestParseStr_1(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				str: "{{.field}}",
			},
			want:    "value",
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				str: "{{.field_not_exist}}",
			},
			want:    "",
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

			got, err := ParseStr(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
