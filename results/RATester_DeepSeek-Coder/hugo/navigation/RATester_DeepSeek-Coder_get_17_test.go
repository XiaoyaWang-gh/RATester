package navigation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_17(t *testing.T) {
	type args struct {
		key       string
		apply     func(m Menu)
		menuLists []Menu
	}
	tests := []struct {
		name  string
		c     *menuCache
		args  args
		want  Menu
		want1 bool
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

			got, got1 := tt.c.get(tt.args.key, tt.args.apply, tt.args.menuLists...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menuCache.get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("menuCache.get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
