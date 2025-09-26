package cssjs

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestResolve_1(t *testing.T) {
	type args struct {
		imp *importResolver
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				imp: &importResolver{
					r:                 nil,
					inPath:            "",
					opts:              InlineImports{},
					contentSeen:       map[string]bool{},
					dependencyManager: nil,
					linemap:           map[int]fileOffset{},
					fs:                nil,
					logger:            nil,
				},
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

			got, err := tt.args.imp.resolve()
			if (err != nil) != tt.wantErr {
				t.Errorf("resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolve() got = %v, want %v", got, tt.want)
			}
		})
	}
}
