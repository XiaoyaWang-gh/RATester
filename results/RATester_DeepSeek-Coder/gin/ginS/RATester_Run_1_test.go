package ginS

// import (
// 	"fmt"
// 	"testing"
// )

// func TestRun_1(t *testing.T) {
// 	type args struct {
// 		addr []string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "TestRun_0",
// 			args: args{
// 				addr: []string{"localhost:8080"},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "TestRun_1",
// 			args: args{
// 				addr: []string{"localhost:8081"},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "TestRun_2",
// 			args: args{
// 				addr: []string{"localhost:8082"},
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			defer func() {
// 				if r := recover(); r != nil {
// 					fmt.Println("Recovered in main", r)
// 				}
// 			}()

// 			if err := Run(tt.args.addr...); (err != nil) != tt.wantErr {
// 				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
