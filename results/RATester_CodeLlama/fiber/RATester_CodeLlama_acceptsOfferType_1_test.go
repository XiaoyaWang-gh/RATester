package fiber

import (
	"fmt"
	"testing"
)

func TestAcceptsOfferType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var spec string
	var offerType string
	var specParams headerParams

	if acceptsOfferType(spec, offerType, specParams) != false {
		t.Errorf("acceptsOfferType() = %v, want %v", acceptsOfferType(spec, offerType, specParams), false)
	}
}
