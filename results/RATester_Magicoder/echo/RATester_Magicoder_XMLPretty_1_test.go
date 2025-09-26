package echo

import (
	"fmt"
	"testing"
)

func TestXMLPretty_1(t *testing.T) {
	type args struct {
		code   int
		i      interface{}
		indent string
	}
	tests := []struct {
		name    string
		c       *context
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

			if err := tt.c.XMLPretty(tt.args.code, tt.args.i, tt.args.indent); (err != nil) != tt.wantErr {
				t.Errorf("context.XMLPretty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
