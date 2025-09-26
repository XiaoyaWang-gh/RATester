package layouts

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddSectionType_1(t *testing.T) {
	type args struct {
		d LayoutDescriptor
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test with empty section",
			args: args{
				d: LayoutDescriptor{
					Section: "",
				},
			},
			want: []string{},
		},
		{
			name: "Test with non-empty section",
			args: args{
				d: LayoutDescriptor{
					Section: "section",
				},
			},
			want: []string{"section"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &layoutBuilder{
				d: tt.args.d,
			}
			l.addSectionType()
			if !reflect.DeepEqual(l.typeVariations, tt.want) {
				t.Errorf("layoutBuilder.addSectionType() = %v, want %v", l.typeVariations, tt.want)
			}
		})
	}
}
