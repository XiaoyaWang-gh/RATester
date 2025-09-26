package config

import (
	"fmt"
	"testing"
)

func TestInt_2(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				key: "field name",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				key: "field name",
			},
			want:    2,
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

			got, err := Int(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
