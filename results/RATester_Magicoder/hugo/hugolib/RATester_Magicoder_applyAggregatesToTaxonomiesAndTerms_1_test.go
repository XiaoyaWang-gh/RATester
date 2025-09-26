package hugolib

import (
	"context"
	"fmt"
	"testing"
)

func TestapplyAggregatesToTaxonomiesAndTerms_1(t *testing.T) {
	type fields struct {
		Site            *Site
		assembleChanges *WhatChanged
		ctx             context.Context
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
				Site:            tt.fields.Site,
				assembleChanges: tt.fields.assembleChanges,
				ctx:             tt.fields.ctx,
			}
			if err := sa.applyAggregatesToTaxonomiesAndTerms(); (err != nil) != tt.wantErr {
				t.Errorf("sitePagesAssembler.applyAggregatesToTaxonomiesAndTerms() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
