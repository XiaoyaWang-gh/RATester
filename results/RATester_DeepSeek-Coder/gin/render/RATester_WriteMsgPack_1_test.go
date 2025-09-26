package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteMsgPack_1(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		obj any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				w:   httptest.NewRecorder(),
				obj: "test",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				w:   httptest.NewRecorder(),
				obj: 123,
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				w:   httptest.NewRecorder(),
				obj: struct{ Name string }{Name: "test"},
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				w:   httptest.NewRecorder(),
				obj: make(chan int),
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

			err := WriteMsgPack(tt.args.w, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteMsgPack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
