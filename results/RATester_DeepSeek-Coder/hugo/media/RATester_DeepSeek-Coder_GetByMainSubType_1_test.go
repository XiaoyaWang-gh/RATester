package media

import (
	"fmt"
	"testing"
)

func TestGetByMainSubType_1(t *testing.T) {
	tests := []struct {
		name      string
		types     Types
		mainType  string
		subType   string
		wantTp    Type
		wantFound bool
	}{
		{
			name: "Test case 1",
			types: Types{
				{MainType: "application", SubType: "json"},
				{MainType: "text", SubType: "html"},
			},
			mainType:  "application",
			subType:   "json",
			wantTp:    Type{MainType: "application", SubType: "json"},
			wantFound: true,
		},
		{
			name: "Test case 2",
			types: Types{
				{MainType: "application", SubType: "json"},
				{MainType: "text", SubType: "html"},
			},
			mainType:  "text",
			subType:   "html",
			wantTp:    Type{MainType: "text", SubType: "html"},
			wantFound: true,
		},
		{
			name: "Test case 3",
			types: Types{
				{MainType: "application", SubType: "json"},
				{MainType: "text", SubType: "html"},
			},
			mainType:  "application",
			subType:   "xml",
			wantTp:    Type{},
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tp, found := tt.types.GetByMainSubType(tt.mainType, tt.subType)
			if tp != tt.wantTp || found != tt.wantFound {
				t.Errorf("GetByMainSubType() = %v, %v; want %v, %v", tp, found, tt.wantTp, tt.wantFound)
			}
		})
	}
}
