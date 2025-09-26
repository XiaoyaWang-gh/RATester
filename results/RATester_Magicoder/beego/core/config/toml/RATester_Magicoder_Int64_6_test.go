package toml

import (
	"fmt"
	"testing"
)

func TestInt64_6(t *testing.T) {
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
				key: "a.b.c",
			},
			want:    123,
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				key: "a.b.d",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test case 3",
			args: args{
				key: "a.b.e",
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

			c := &configContainer{}
			got, err := c.Int64(tt.args.key)
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
