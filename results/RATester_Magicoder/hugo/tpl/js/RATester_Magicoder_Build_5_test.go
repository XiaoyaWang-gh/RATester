package js

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestBuild_5(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name    string
		ns      *Namespace
		args    args
		want    resource.Resource
		wantErr bool
	}{
		{
			name: "Test case 1",
			ns:   &Namespace{},
			args: args{
				args: []any{"test", "test"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test case 2",
			ns:   &Namespace{},
			args: args{
				args: []any{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test case 3",
			ns:   &Namespace{},
			args: args{
				args: []any{1, 2, 3},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test case 4",
			ns:   &Namespace{},
			args: args{
				args: []any{"test"},
			},
			want:    nil,
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

			got, err := tt.ns.Build(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
