package conn

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLenBytes_1(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		args    args
		wantB   []byte
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				buf: []byte("test"),
			},
			wantB:   append([]byte{0, 0, 0, 4}, []byte("test")...),
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				buf: []byte(""),
			},
			wantB:   append([]byte{0, 0, 0, 0}, []byte("")...),
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				buf: nil,
			},
			wantB:   append([]byte{0, 0, 0, 0}, []byte("")...),
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

			gotB, err := GetLenBytes(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLenBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("GetLenBytes() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
