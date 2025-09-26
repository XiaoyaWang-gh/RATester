package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestDeepCopyObject_1(t *testing.T) {
	tests := []struct {
		name string
		in   *TraefikServiceList
		want runtime.Object
	}{
		{
			name: "case 1",
			in: &TraefikServiceList{
				TypeMeta: metav1.TypeMeta{
					Kind:       "TraefikServiceList",
					APIVersion: "v1alpha1",
				},
				ListMeta: metav1.ListMeta{
					ResourceVersion: "1",
				},
				Items: []TraefikService{
					{
						TypeMeta: metav1.TypeMeta{
							Kind:       "TraefikService",
							APIVersion: "v1alpha1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test",
							Namespace: "default",
						},
					},
				},
			},
			want: &TraefikServiceList{
				TypeMeta: metav1.TypeMeta{
					Kind:       "TraefikServiceList",
					APIVersion: "v1alpha1",
				},
				ListMeta: metav1.ListMeta{
					ResourceVersion: "1",
				},
				Items: []TraefikService{
					{
						TypeMeta: metav1.TypeMeta{
							Kind:       "TraefikService",
							APIVersion: "v1alpha1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test",
							Namespace: "default",
						},
					},
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

			got := tt.in.DeepCopyObject()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
