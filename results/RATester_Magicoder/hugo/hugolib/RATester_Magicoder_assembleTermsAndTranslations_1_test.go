package hugolib

import (
	"fmt"
	"testing"
)

func TestassembleTermsAndTranslations_1(t *testing.T) {
	type fields struct {
		Site *Site
	}
	tests := []struct {
		name    string
		fields  fields
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

			sa := &sitePagesAssembler{
				Site: tt.fields.Site,
			}
			if err := sa.assembleTermsAndTranslations(); (err != nil) != tt.wantErr {
				t.Errorf("sitePagesAssembler.assembleTermsAndTranslations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
