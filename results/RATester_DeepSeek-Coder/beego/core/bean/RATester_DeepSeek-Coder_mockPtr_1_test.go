package bean

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMockPtr_1(t *testing.T) {
	type args struct {
		pvv reflect.Value
		ptt reflect.Type
	}
	tests := []struct {
		name    string
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

			err := mockPtr(tt.args.pvv, tt.args.ptt)
			if (err != nil) != tt.wantErr {
				t.Errorf("mockPtr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
