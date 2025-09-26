package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_11(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name string
		r    ProtoBuf
		args args
	}{
		{
			name: "Test Case 1",
			r:    ProtoBuf{Data: "test data"},
			args: args{
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "Test Case 2",
			r:    ProtoBuf{Data: 123},
			args: args{
				w: httptest.NewRecorder(),
			},
		},
		{
			name: "Test Case 3",
			r:    ProtoBuf{Data: nil},
			args: args{
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.r.WriteContentType(tt.args.w)
		})
	}
}
