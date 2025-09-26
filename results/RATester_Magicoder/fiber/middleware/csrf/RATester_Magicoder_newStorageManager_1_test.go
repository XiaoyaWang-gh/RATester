package csrf

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestnewStorageManager_1(t *testing.T) {
	type args struct {
		storage fiber.Storage
	}
	tests := []struct {
		name string
		args args
		want *storageManager
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

			if got := newStorageManager(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStorageManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
