package migration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddIndex_1(t *testing.T) {
	type args struct {
		index *Index
	}
	tests := []struct {
		name string
		m    *Migration
		args args
		want *Migration
	}{
		{
			name: "TestAddIndex",
			m:    &Migration{},
			args: args{
				index: &Index{
					Name: "test_index",
				},
			},
			want: &Migration{
				Indexes: []*Index{
					{
						Name: "test_index",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.AddIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Migration.AddIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
