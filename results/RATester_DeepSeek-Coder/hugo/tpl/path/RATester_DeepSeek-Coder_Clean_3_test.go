package path

import (
	"fmt"
	"testing"
)

func TestClean_3(t *testing.T) {
	type args struct {
		path any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestClean/Success",
			args: args{
				path: "/foo/bar",
			},
			want:    "/foo/bar",
			wantErr: false,
		},
		{
			name: "TestClean/Error",
			args: args{
				path: 123,
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

			ns := &Namespace{}
			got, err := ns.Clean(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Clean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.Clean() = %v, want %v", got, tt.want)
			}
		})
	}
}
