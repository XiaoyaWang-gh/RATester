package langs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewLanguage_1(t *testing.T) {
	type args struct {
		lang                   string
		defaultContentLanguage string
		timeZone               string
		languageConfig         LanguageConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *Language
		wantErr bool
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

			got, err := NewLanguage(tt.args.lang, tt.args.defaultContentLanguage, tt.args.timeZone, tt.args.languageConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLanguage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}
