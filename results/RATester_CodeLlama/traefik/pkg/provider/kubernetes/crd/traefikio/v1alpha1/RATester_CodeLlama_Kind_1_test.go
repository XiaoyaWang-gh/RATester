package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	kschema "k8s.io/apimachinery/pkg/runtime/schema"
)

func TestKind_1(t *testing.T) {
	tests := []struct {
		name string
		kind string
		want kschema.GroupKind
	}{
		{
			name: "test case 1",
			kind: "test case 1",
			want: kschema.GroupKind{
				Group: "",
				Kind:  "test case 1",
			},
		},
		{
			name: "test case 2",
			kind: "test case 2",
			want: kschema.GroupKind{
				Group: "",
				Kind:  "test case 2",
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

			if got := Kind(tt.kind); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}
