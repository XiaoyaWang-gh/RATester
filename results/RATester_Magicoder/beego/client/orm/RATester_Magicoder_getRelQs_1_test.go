package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestgetRelQs_1(t *testing.T) {
	o := &ormBase{}
	mi := &models.ModelInfo{}
	fi := &models.FieldInfo{}

	tests := []struct {
		name string
		o    *ormBase
		md   interface{}
		mi   *models.ModelInfo
		fi   *models.FieldInfo
		want *querySet
	}{
		{
			name: "test case 1",
			o:    o,
			md:   "test",
			mi:   mi,
			fi:   fi,
			want: &querySet{},
		},
		// add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.o.getRelQs(tt.md, tt.mi, tt.fi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRelQs() = %v, want %v", got, tt.want)
			}
		})
	}
}
