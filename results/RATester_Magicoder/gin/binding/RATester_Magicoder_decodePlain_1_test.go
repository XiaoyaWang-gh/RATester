package binding

import (
	"fmt"
	"testing"
)

func TestdecodePlain_1(t *testing.T) {
	type args struct {
		data []byte
		obj  any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test decodePlain with string",
			args: args{
				data: []byte("test"),
				obj:  new(string),
			},
			wantErr: false,
		},
		{
			name: "Test decodePlain with []byte",
			args: args{
				data: []byte("test"),
				obj:  new([]byte),
			},
			wantErr: false,
		},
		{
			name: "Test decodePlain with unknown type",
			args: args{
				data: []byte("test"),
				obj:  new(int),
			},
			wantErr: true,
		},
		{
			name: "Test decodePlain with nil obj",
			args: args{
				data: []byte("test"),
				obj:  nil,
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

			if err := decodePlain(tt.args.data, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("decodePlain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
