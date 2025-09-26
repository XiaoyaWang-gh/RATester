package pagemeta

import (
	"fmt"
	"testing"
)

func TestNewChainedFrontMatterFieldHandler_1(t *testing.T) {
	type args struct {
		handlers []frontMatterFieldHandler
	}
	tests := []struct {
		name string
		f    FrontMatterHandler
		args args
		want frontMatterFieldHandler
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

			got := tt.f.newChainedFrontMatterFieldHandler(tt.args.handlers...)
			if got == nil {
				t.Errorf("newChainedFrontMatterFieldHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
