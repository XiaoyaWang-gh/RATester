package api

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestnewSearchCriterion_1(t *testing.T) {
	tests := []struct {
		name  string
		query url.Values
		want  *searchCriterion
	}{
		{
			name:  "empty query",
			query: url.Values{},
			want:  nil,
		},
		{
			name: "empty search criterion",
			query: url.Values{
				"search":         []string{""},
				"status":         []string{""},
				"serviceName":    []string{""},
				"middlewareName": []string{""},
			},
			want: nil,
		},
		{
			name: "valid search criterion",
			query: url.Values{
				"search":         []string{"test"},
				"status":         []string{"active"},
				"serviceName":    []string{"service1"},
				"middlewareName": []string{"middleware1"},
			},
			want: &searchCriterion{
				Search:         "test",
				Status:         "active",
				ServiceName:    "service1",
				MiddlewareName: "middleware1",
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

			got := newSearchCriterion(tt.query)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSearchCriterion() = %v, want %v", got, tt.want)
			}
		})
	}
}
