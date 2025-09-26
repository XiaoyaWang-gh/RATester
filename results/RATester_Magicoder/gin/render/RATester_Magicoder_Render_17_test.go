package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_17(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		r       MsgPack
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			r: MsgPack{
				Data: "test data",
			},
			args: args{
				w: httptest.NewRecorder(),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			r: MsgPack{
				Data: nil,
			},
			args: args{
				w: httptest.NewRecorder(),
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			r: MsgPack{
				Data: make(chan int),
			},
			args: args{
				w: httptest.NewRecorder(),
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

			if err := tt.r.Render(tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("MsgPack.Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
