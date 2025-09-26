package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRenderContent_1(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				in: []byte("Hello, {{.Name}}!"),
			},
			wantOut: []byte("Hello, John!"),
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				in: []byte("Today is {{.Day}}."),
			},
			wantOut: []byte("Today is Monday."),
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				in: []byte("{{.Invalid}}"),
			},
			wantOut: []byte(""),
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

			out, err := RenderContent(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(out, tt.wantOut) {
				t.Errorf("RenderContent() = %v, want %v", out, tt.wantOut)
			}
		})
	}
}
