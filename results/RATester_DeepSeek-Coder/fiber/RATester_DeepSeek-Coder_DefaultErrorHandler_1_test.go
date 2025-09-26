package fiber

import (
	"fmt"
	"testing"
)

func TestDefaultErrorHandler_1(t *testing.T) {
	type args struct {
		c   Ctx
		err error
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

			if err := DefaultErrorHandler(tt.args.c, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("DefaultErrorHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
