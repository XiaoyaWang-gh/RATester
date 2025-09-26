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
			name: "Test Case 1",
			kind: "Pod",
			want: kschema.GroupKind{
				Group: "",
				Kind:  "Pod",
			},
		},
		{
			name: "Test Case 2",
			kind: "Service",
			want: kschema.GroupKind{
				Group: "",
				Kind:  "Service",
			},
		},
		{
			name: "Test Case 3",
			kind: "ReplicaSet",
			want: kschema.GroupKind{
				Group: "apps",
				Kind:  "ReplicaSet",
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
