package middleware

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestValuesFromParam_1(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name    string
		args    args
		want    ValuesExtractor
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				param: "test",
			},
			want: func(c echo.Context) ([]string, error) {
				result := make([]string, 0)
				paramVales := c.ParamValues()
				for i, p := range c.ParamNames() {
					if "test" == p {
						result = append(result, paramVales[i])
						if i >= extractorLimit-1 {
							break
						}
					}
				}
				if len(result) == 0 {
					return nil, errParamExtractorValueMissing
				}
				return result, nil
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := valuesFromParam(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("valuesFromParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
