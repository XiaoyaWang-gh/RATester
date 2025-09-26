package converter

import (
	"fmt"
	"testing"
)

func TestConvert_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var ctx RenderContext
	var nopConverter nopConverter
	var result ResultRender
	var err error

	result, err = nopConverter.Convert(ctx)
	if err != nil {
		t.Errorf("Convert() error = %v", err)
		return
	}

	if result == nil {
		t.Errorf("Convert() result = %v", result)
		return
	}

	if result.Bytes() == nil {
		t.Errorf("Convert() result.Bytes() = %v", result.Bytes())
		return
	}
}
