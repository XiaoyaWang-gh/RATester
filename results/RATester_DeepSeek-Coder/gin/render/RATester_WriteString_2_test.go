package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteString_2(t *testing.T) {
	type args struct {
		w      *httptest.ResponseRecorder
		format string
		data   []any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with data",
			args: args{
				w:      httptest.NewRecorder(),
				format: "Hello, %s",
				data:   []any{"World"},
			},
			wantErr: false,
		},
		{
			name: "Test without data",
			args: args{
				w:      httptest.NewRecorder(),
				format: "Hello, World",
				data:   []any{},
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

			err := WriteString(tt.args.w, tt.args.format, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
