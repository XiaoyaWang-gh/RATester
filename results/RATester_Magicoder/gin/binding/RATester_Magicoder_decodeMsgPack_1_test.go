package binding

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestdecodeMsgPack_1(t *testing.T) {
	type args struct {
		r   io.Reader
		obj any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test decodeMsgPack with valid input",
			args: args{
				r: strings.NewReader(`{"field_name": "value"}`),
				obj: struct {
					FieldName string `json:"field_name"`
				}{},
			},
			wantErr: false,
		},
		{
			name: "Test decodeMsgPack with invalid input",
			args: args{
				r: strings.NewReader(`{"field_name": "value"`),
				obj: struct {
					FieldName string `json:"field_name"`
				}{},
			},
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

			if err := decodeMsgPack(tt.args.r, &tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("decodeMsgPack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
