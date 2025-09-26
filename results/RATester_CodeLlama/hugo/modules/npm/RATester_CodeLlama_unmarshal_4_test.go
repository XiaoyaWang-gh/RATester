package npm

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestUnmarshal_4(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		b       *packageBuilder
		args    args
		want    map[string]any
		wantErr bool
	}{
		{
			name: "test",
			b: &packageBuilder{
				originalPackageJSON: map[string]any{
					"name": "test",
				},
			},
			args: args{
				r: strings.NewReader(`{"name": "test"}`),
			},
			want: map[string]any{
				"name": "test",
			},
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

			if got := tt.b.unmarshal(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("packageBuilder.unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
