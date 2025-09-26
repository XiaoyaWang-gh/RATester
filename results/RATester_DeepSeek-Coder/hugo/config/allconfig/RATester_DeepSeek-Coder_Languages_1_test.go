package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestLanguages_1(t *testing.T) {
	type fields struct {
		m *Configs
	}
	tests := []struct {
		name   string
		fields fields
		want   langs.Languages
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

			c := ConfigLanguage{
				m: tt.fields.m,
			}
			if got := c.Languages(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Languages() = %v, want %v", got, tt.want)
			}
		})
	}
}
