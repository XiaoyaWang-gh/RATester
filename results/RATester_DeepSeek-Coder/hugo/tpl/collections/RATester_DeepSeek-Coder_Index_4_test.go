package collections

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/tpl/compare"
)

func TestIndex_4(t *testing.T) {
	ns := &Namespace{
		loc:      time.UTC,
		sortComp: &compare.Namespace{},
		deps:     &deps.Deps{},
	}

	tests := []struct {
		name    string
		item    any
		args    []any
		want    any
		wantErr bool
	}{
		{
			name: "test1",
			item: []int{1, 2, 3, 4, 5},
			args: []any{2},
			want: 2,
		},
		{
			name: "test2",
			item: []string{"a", "b", "c", "d", "e"},
			args: []any{"c"},
			want: "c",
		},
		{
			name:    "test3",
			item:    []int{1, 2, 3, 4, 5},
			args:    []any{6},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.Index(tt.item, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}
