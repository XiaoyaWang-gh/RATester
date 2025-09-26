package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestGetDependencyManager_1(t *testing.T) {
	type fields struct {
		sd ResourceSourceDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   identity.Manager
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

			l := &genericResource{
				sd: tt.fields.sd,
			}
			if got := l.GetDependencyManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genericResource.GetDependencyManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
