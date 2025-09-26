package middleware

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt"
)

func TestDefaultKeyFunc_1(t *testing.T) {
	type args struct {
		config *JWTConfig
		token  *jwt.Token
	}
	tests := []struct {
		name    string
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

			got, err := tt.args.config.defaultKeyFunc(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultKeyFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultKeyFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
