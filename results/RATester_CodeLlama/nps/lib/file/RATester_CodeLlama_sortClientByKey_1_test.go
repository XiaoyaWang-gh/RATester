package file

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestSortClientByKey_1(t *testing.T) {
	type args struct {
		m       sync.Map
		sortKey string
		order   string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				m:       sync.Map{},
				sortKey: "id",
				order:   "asc",
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := sortClientByKey(tt.args.m, tt.args.sortKey, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortClientByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
