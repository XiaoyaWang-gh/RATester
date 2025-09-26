package publisher

import (
	"fmt"
	"testing"
)

func TestCreateTransformerChain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := DestinationPublisher{}
	f := Descriptor{}

	transformers := p.createTransformerChain(f)

	if len(transformers) != 0 {
		t.Errorf("Expected empty transformer chain, got %v", transformers)
	}
}
