package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParams_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{
		pageMetaParams: &pageMetaParams{
			paramsOriginal: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
		},
	}

	if got := p.Params(); !reflect.DeepEqual(got, p.pageConfig.Params) {
		t.Errorf("Params() = %v, want %v", got, p.pageConfig.Params)
	}
}
