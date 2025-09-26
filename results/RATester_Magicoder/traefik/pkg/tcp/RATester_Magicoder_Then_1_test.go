package tcp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThen_1(t *testing.T) {
	type args struct {
		h Handler
	}
	tests := []struct {
		name    string
		c       Chain
		args    args
		want    Handler
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

			got, err := tt.c.Then(tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain.Then() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chain.Then() = %v, want %v", got, tt.want)
			}
		})
	}
}
