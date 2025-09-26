package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
	"github.com/beego/beego/v2/client/orm/internal/utils"
)

func TestGetColumnDefault_1(t *testing.T) {
	type testCase struct {
		name     string
		fi       *models.FieldInfo
		expected string
	}

	testCases := []testCase{
		{
			name: "Test with Rel field",
			fi: &models.FieldInfo{
				Rel: true,
			},
			expected: "",
		},
		{
			name: "Test with Reverse field",
			fi: &models.FieldInfo{
				Reverse: true,
			},
			expected: "",
		},
		{
			name: "Test with non-relevant field type",
			fi: &models.FieldInfo{
				FieldType: models.TypeTextField,
			},
			expected: "",
		},
		{
			name: "Test with relevant field type",
			fi: &models.FieldInfo{
				FieldType:  models.TypeIntegerField,
				ColDefault: true,
				Initial:    utils.StrTo("10"),
			},
			expected: " DEFAULT 10 ",
		},
		{
			name: "Test with relevant field type and no initial value",
			fi: &models.FieldInfo{
				FieldType:  models.TypeIntegerField,
				ColDefault: true,
			},
			expected: " DEFAULT 0 ",
		},
		{
			name: "Test with relevant field type and null",
			fi: &models.FieldInfo{
				FieldType: models.TypeIntegerField,
				Null:      true,
			},
			expected: " DEFAULT 0 ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := getColumnDefault(tc.fi)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
