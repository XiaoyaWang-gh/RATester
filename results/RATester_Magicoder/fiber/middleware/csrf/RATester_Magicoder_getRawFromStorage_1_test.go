package csrf

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestgetRawFromStorage_1(t *testing.T) {
	type args struct {
		c              fiber.Ctx
		token          string
		cfg            Config
		sessionManager *sessionManager
		storageManager *storageManager
	}
	tests := []struct {
		name string
		args args
		want []byte
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

			if got := getRawFromStorage(tt.args.c, tt.args.token, tt.args.cfg, tt.args.sessionManager, tt.args.storageManager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRawFromStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}
