package pagemeta

import (
	"fmt"
	"testing"
)

func TestCreateContentAdapterDatesHandler_1(t *testing.T) {
	tests := []struct {
		name    string
		fmcfg   FrontmatterConfig
		desc    FrontMatterDescriptor
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

			f := FrontMatterHandler{}
			got, err := f.createContentAdapterDatesHandler(tt.fmcfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("FrontMatterHandler.createContentAdapterDatesHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := got(&tt.desc); (err != nil) != tt.wantErr {
				t.Errorf("FrontMatterHandler.createContentAdapterDatesHandler() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
