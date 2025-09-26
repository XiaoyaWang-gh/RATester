package cache

import (
	"fmt"
	"testing"
)

func Testexists_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				path: "testdata/test.txt",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				path: "testdata/nonexistent.txt",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				path: "",
			},
			want:    false,
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

			got, err := exists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
