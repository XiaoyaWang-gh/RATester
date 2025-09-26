package binding

import (
	"fmt"
	"testing"
)

func TestDecodePlain_1(t *testing.T) {
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
			name: "case1",
			args: args{
				data: []byte("test"),
				obj:  "test",
			},
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				data: []byte("test"),
				obj:  []byte("test"),
			},
			wantErr: false,
		},
		{
			name: "case3",
			args: args{
				data: []byte("test"),
				obj:  nil,
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

			if err := decodePlain(tt.args.data, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("decodePlain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
