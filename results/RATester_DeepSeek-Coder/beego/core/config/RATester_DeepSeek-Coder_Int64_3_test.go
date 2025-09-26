package config

import (
	"fmt"
	"testing"
)

func TestInt64_3(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				key: "test_key",
			},
			want:    123,
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				key: "non_existing_key",
			},
			want:    0,
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

			got, err := Int64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}
