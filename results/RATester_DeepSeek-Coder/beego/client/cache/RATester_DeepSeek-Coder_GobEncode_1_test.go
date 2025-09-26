package cache

import (
	"fmt"
	"testing"
)

func TestGobEncode_1(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestGobEncode_Success",
			args: args{
				data: "test",
			},
			wantErr: false,
		},
		{
			name: "TestGobEncode_Fail",
			args: args{
				data: make(chan int),
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

			_, err := GobEncode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
