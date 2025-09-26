package csrf

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v3/middleware/session"
)

func TestNewSessionManager_1(t *testing.T) {
	type args struct {
		s *session.Store
		k string
	}
	tests := []struct {
		name string
		args args
		want *sessionManager
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

			if got := newSessionManager(tt.args.s, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSessionManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
