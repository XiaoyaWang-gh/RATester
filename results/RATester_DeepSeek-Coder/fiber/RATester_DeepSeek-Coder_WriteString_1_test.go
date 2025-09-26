package fiber

import (
	"fmt"
	"testing"
)

func TestWriteString_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		c       *DefaultCtx
		args    args
		want    int
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

			got, err := tt.c.WriteString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultCtx.WriteString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DefaultCtx.WriteString() = %v, want %v", got, tt.want)
			}
		})
	}
}
