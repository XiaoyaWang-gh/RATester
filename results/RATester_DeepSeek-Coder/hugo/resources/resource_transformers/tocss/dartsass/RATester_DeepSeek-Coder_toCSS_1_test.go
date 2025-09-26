package dartsass

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/bep/godartsass/v2"
)

func TestToCSS_1(t *testing.T) {
	type args struct {
		args godartsass.Args
		src  io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    godartsass.Result
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

			c := &Client{}
			got, err := c.toCSS(tt.args.args, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.toCSS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.toCSS() = %v, want %v", got, tt.want)
			}
		})
	}
}
