package middleware

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPush_1(t *testing.T) {
	type args struct {
		target string
		opts   *http.PushOptions
	}
	tests := []struct {
		name    string
		w       *gzipResponseWriter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.w.Push(tt.args.target, tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("gzipResponseWriter.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
