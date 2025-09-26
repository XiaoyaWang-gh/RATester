package middleware

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestDefaultParseToken_1(t *testing.T) {
	type args struct {
		auth string
		c    echo.Context
	}
	tests := []struct {
		name    string
		config  *JWTConfig
		args    args
		want    interface{}
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

			got, err := tt.config.defaultParseToken(tt.args.auth, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultParseToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
