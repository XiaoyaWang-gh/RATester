package bean

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewTagAutoWireBeanFactory_1(t *testing.T) {
	type args struct {
		adapters       map[string]TypeAdapter
		fieldTagParser func(field reflect.StructField) *FieldMetadata
	}
	tests := []struct {
		name string
		args args
		want *TagAutoWireBeanFactory
	}{
		{
			name: "test1",
			args: args{
				adapters: map[string]TypeAdapter{
					"Time": &TimeTypeAdapter{Layout: "2006-01-02 15:04:05"},
				},
				fieldTagParser: func(field reflect.StructField) *FieldMetadata {
					return &FieldMetadata{
						DftValue: field.Tag.Get(DefaultValueTagKey),
					}
				},
			},
			want: &TagAutoWireBeanFactory{
				Adapters: map[string]TypeAdapter{
					"Time": &TimeTypeAdapter{Layout: "2006-01-02 15:04:05"},
				},
				FieldTagParser: func(field reflect.StructField) *FieldMetadata {
					return &FieldMetadata{
						DftValue: field.Tag.Get(DefaultValueTagKey),
					}
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

			if got := NewTagAutoWireBeanFactory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagAutoWireBeanFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
