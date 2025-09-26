package adaptor

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestConvertRequest_1(t *testing.T) {
	type args struct {
		c         fiber.Ctx
		forServer bool
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
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

			got, err := ConvertRequest(tt.args.c, tt.args.forServer)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
