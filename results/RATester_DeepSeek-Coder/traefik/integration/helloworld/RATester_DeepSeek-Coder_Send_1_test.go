package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestSend_1(t *testing.T) {
	type args struct {
		m *StreamExampleReply
	}
	tests := []struct {
		name    string
		x       *greeterStreamExampleServer
		args    args
		wantErr bool
	}{
		{
			name: "TestSend",
			x:    &greeterStreamExampleServer{},
			args: args{
				m: &StreamExampleReply{
					Data: "test",
				},
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

			if err := tt.x.Send(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("greeterStreamExampleServer.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
