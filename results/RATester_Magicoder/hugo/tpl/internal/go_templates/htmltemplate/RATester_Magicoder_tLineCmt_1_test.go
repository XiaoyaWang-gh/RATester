package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttLineCmt_1(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context
		input   []byte
		want    context
		wantN   int
		wantErr bool
	}{
		{
			name: "test case 1",
			ctx: context{
				state: stateJSLineCmt,
			},
			input: []byte("// this is a line comment\n"),
			want: context{
				state: stateJS,
			},
			wantN:   24,
			wantErr: false,
		},
		{
			name: "test case 2",
			ctx: context{
				state: stateJSLineCmt,
			},
			input: []byte("/* this is a block comment\n"),
			want: context{
				state: stateJS,
			},
			wantN:   24,
			wantErr: false,
		},
		{
			name: "test case 3",
			ctx: context{
				state: stateJSLineCmt,
			},
			input: []byte("this is not a comment\n"),
			want: context{
				state: stateJS,
			},
			wantN:   19,
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

			got, gotN := tLineCmt(tt.ctx, tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tLineCmt() got = %v, want %v", got, tt.want)
			}
			if gotN != tt.wantN {
				t.Errorf("tLineCmt() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
