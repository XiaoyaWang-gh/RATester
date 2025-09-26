package goldmark

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/util"
)

func TestPut_1(t *testing.T) {
	type fields struct {
		idType string
		vals   map[string]struct{}
	}
	type args struct {
		value []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestPut",
			fields: fields{
				idType: "test",
				vals:   make(map[string]struct{}),
			},
			args: args{
				value: []byte("test"),
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

			ids := &idFactory{
				idType: tt.fields.idType,
				vals:   tt.fields.vals,
			}
			ids.Put(tt.args.value)
			if _, ok := ids.vals[util.BytesToReadOnlyString(tt.args.value)]; !ok {
				t.Errorf("Put() error, value not found in map")
			}
		})
	}
}
