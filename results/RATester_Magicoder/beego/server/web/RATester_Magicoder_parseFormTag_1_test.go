package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestparseFormTag_1(t *testing.T) {
	type args struct {
		fieldT reflect.StructField
	}
	tests := []struct {
		name         string
		args         args
		wantLabel    string
		wantName     string
		wantFType    string
		wantId       string
		wantClass    string
		wantIgnored  bool
		wantRequired bool
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

			gotLabel, gotName, gotFType, gotId, gotClass, gotIgnored, gotRequired := parseFormTag(tt.args.fieldT)
			if gotLabel != tt.wantLabel {
				t.Errorf("parseFormTag() gotLabel = %v, want %v", gotLabel, tt.wantLabel)
			}
			if gotName != tt.wantName {
				t.Errorf("parseFormTag() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotFType != tt.wantFType {
				t.Errorf("parseFormTag() gotFType = %v, want %v", gotFType, tt.wantFType)
			}
			if gotId != tt.wantId {
				t.Errorf("parseFormTag() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotClass != tt.wantClass {
				t.Errorf("parseFormTag() gotClass = %v, want %v", gotClass, tt.wantClass)
			}
			if gotIgnored != tt.wantIgnored {
				t.Errorf("parseFormTag() gotIgnored = %v, want %v", gotIgnored, tt.wantIgnored)
			}
			if gotRequired != tt.wantRequired {
				t.Errorf("parseFormTag() gotRequired = %v, want %v", gotRequired, tt.wantRequired)
			}
		})
	}
}
