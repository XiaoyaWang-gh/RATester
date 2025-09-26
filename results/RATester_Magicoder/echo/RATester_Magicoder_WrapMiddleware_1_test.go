package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWrapMiddleware_1(t *testing.T) {
	type args struct {
		m    func(http.Handler) http.Handler
		next HandlerFunc
		c    Context
	}
	tests := []struct {
		name    string
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

			if err := WrapMiddleware(tt.args.m)(tt.args.next)(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("WrapMiddleware() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
