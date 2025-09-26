package hugofs

import (
	"fmt"
	"testing"
)

func Testfilename_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		r    RootMapping
		args args
		want string
	}{
		{
			name: "Test case 1",
			r:    RootMapping{From: "test"},
			args: args{name: "test1"},
			want: "test1",
		},
		{
			name: "Test case 2",
			r:    RootMapping{From: "test"},
			args: args{name: ""},
			want: "test",
		},
		{
			name: "Test case 3",
			r:    RootMapping{From: "test", To: "test2"},
			args: args{name: "test1"},
			want: "test2/test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.r.filename(tt.args.name); got != tt.want {
				t.Errorf("filename() = %v, want %v", got, tt.want)
			}
		})
	}
}
