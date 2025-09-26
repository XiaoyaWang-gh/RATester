package helloworld

import (
	fmt "fmt"
	"testing"

	grpc "google.golang.org/grpc"
)

func Test_Greeter_StreamExample_Handler_1(t *testing.T) {

	type args struct {
		srv    interface{}
		stream grpc.ServerStream
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test_Greeter_StreamExample_Handler_0",
			args: args{
				srv:    nil,
				stream: nil,
			},
			wantErr: true,
		},
		{
			name: "Test_Greeter_StreamExample_Handler_1",
			args: args{
				srv:    nil,
				stream: nil,
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

			if err := _Greeter_StreamExample_Handler(tt.args.srv, tt.args.stream); (err != nil) != tt.wantErr {
				t.Errorf("_Greeter_StreamExample_Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
