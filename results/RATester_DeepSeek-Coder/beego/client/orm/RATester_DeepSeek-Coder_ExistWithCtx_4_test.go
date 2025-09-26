package orm

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestExistWithCtx_4(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testCases := []struct {
		name     string
		ctx      context.Context
		md       interface{}
		expected bool
	}{
		{
			name:     "ExistWithCtx_Success",
			ctx:      ctx,
			md:       "test",
			expected: true,
		},
		{
			name:     "ExistWithCtx_Fail",
			ctx:      ctx,
			md:       "not_exist",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			qm := &queryM2M{
				md: tc.md,
				fi: &models.FieldInfo{
					ReverseFieldInfo: &models.FieldInfo{
						Name: "reverse_field_name",
					},
					ReverseFieldInfoTwo: &models.FieldInfo{
						Name: "reverse_field_name_two",
					},
				},
				qs: &querySet{
					cond: &Condition{
						params: []condValue{
							{
								exprs: []string{"reverse_field_name"},
								args:  []interface{}{tc.md},
							},
							{
								exprs: []string{"reverse_field_name_two"},
								args:  []interface{}{tc.md},
							},
						},
					},
				},
			}

			result := qm.ExistWithCtx(tc.ctx, tc.md)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
