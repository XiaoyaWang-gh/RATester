package hashing

import (
	"fmt"
	"testing"
)

func TestXXHashFromString_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				s: "test",
			},
			want:    0x24bed08144,
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				s: "test2",
			},
			want:    0x24bed08145,
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				s: "test3",
			},
			want:    0x24bed08146,
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				s: "test4",
			},
			want:    0x24bed08147,
			wantErr: false,
		},
		{
			name: "Test case 5",
			args: args{
				s: "test5",
			},
			want:    0x24bed08148,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := XXHashFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("XXHashFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("XXHashFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
