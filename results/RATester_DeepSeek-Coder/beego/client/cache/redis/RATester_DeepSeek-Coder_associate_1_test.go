package redis

import (
	"fmt"
	"testing"
)

func TestAssociate_1(t *testing.T) {
	type args struct {
		originKey interface{}
	}
	tests := []struct {
		name string
		rc   *Cache
		args args
		want string
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

			if got := tt.rc.associate(tt.args.originKey); got != tt.want {
				t.Errorf("associate() = %v, want %v", got, tt.want)
			}
		})
	}
}
