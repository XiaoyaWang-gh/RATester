package try

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHasCn_1(t *testing.T) {
	type args struct {
		cn string
	}
	tests := []struct {
		name string
		args args
		want ResponseCondition
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

			if got := HasCn(tt.args.cn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasCn() = %v, want %v", got, tt.want)
			}
		})
	}
}
