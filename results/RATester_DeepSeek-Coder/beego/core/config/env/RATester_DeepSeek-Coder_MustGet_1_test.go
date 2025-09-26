package env

import (
	"fmt"
	"testing"
)

func TestMustGet_1(t *testing.T) {
	type args struct {
		key string
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
				key: "TEST_KEY_1",
			},
			want:    "TEST_VALUE_1",
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				key: "TEST_KEY_2",
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

			got, err := MustGet(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("MustGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MustGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
