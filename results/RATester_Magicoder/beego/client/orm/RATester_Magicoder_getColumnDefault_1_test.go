package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
	"github.com/beego/beego/v2/client/orm/internal/utils"
)

func TestgetColumnDefault_1(t *testing.T) {
	tests := []struct {
		name string
		fi   *models.FieldInfo
		want string
	}{
		{
			name: "Test case 1",
			fi: &models.FieldInfo{
				Rel:        false,
				Reverse:    false,
				FieldType:  models.TypeBitField,
				ColDefault: true,
				Initial:    utils.StrTo("1"),
			},
			want: " DEFAULT 1 ",
		},
		{
			name: "Test case 2",
			fi: &models.FieldInfo{
				Rel:        false,
				Reverse:    false,
				FieldType:  models.TypeBitField,
				ColDefault: false,
				Null:       true,
			},
			want: " DEFAULT 0 ",
		},
		{
			name: "Test case 3",
			fi: &models.FieldInfo{
				Rel:        false,
				Reverse:    false,
				FieldType:  models.TypeBitField,
				ColDefault: false,
				Null:       false,
			},
			want: "",
		},
		{
			name: "Test case 4",
			fi: &models.FieldInfo{
				Rel:     true,
				Reverse: false,
			},
			want: "",
		},
		{
			name: "Test case 5",
			fi: &models.FieldInfo{
				Rel:     false,
				Reverse: true,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getColumnDefault(tt.fi); got != tt.want {
				t.Errorf("getColumnDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
