package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLookup_1(t *testing.T) {
	type fields struct {
		nameSpace *nameSpace
	}
	type args struct {
		name string
	}

	ns := &nameSpace{
		set: map[string]*Template{
			"test": &Template{},
		},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Template
	}{
		{
			name: "TestLookup_ExistingTemplate",
			fields: fields{
				nameSpace: ns,
			},
			args: args{
				name: "test",
			},
			want: &Template{},
		},
		{
			name: "TestLookup_NonExistingTemplate",
			fields: fields{
				nameSpace: ns,
			},
			args: args{
				name: "nonexisting",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tmpl := &Template{
				nameSpace: tt.fields.nameSpace,
			}
			if got := tmpl.Lookup(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
