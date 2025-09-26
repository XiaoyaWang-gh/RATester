package middleware

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateExtractors_1(t *testing.T) {
	type args struct {
		lookups    string
		authScheme string
	}
	tests := []struct {
		name    string
		args    args
		want    []ValuesExtractor
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				lookups:    "query:q,param:p,cookie:c,form:f,header:h",
				authScheme: "Bearer",
			},
			want: []ValuesExtractor{
				valuesFromQuery("q"),
				valuesFromParam("p"),
				valuesFromCookie("c"),
				valuesFromForm("f"),
				valuesFromHeader("h", "Bearer "),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				lookups:    "",
				authScheme: "",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				lookups:    "invalid",
				authScheme: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := createExtractors(tt.args.lookups, tt.args.authScheme)
			if (err != nil) != tt.wantErr {
				t.Errorf("createExtractors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createExtractors() = %v, want %v", got, tt.want)
			}
		})
	}
}
