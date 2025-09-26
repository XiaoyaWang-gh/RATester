package urlreplacers

import (
	"fmt"
	"testing"
)

func TestConsumeQuote_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lexer := &absurllexer{
		content: []byte("test content"),
		quotes:  [][]byte{[]byte("quote")},
	}

	result := lexer.consumeQuote()
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}

	lexer.content = append(lexer.content, []byte("quote")...)
	result = lexer.consumeQuote()
	if string(result) != "quote" {
		t.Errorf("Expected 'quote', got %s", result)
	}
	if lexer.pos != len("quote") {
		t.Errorf("Expected pos to be %d, got %d", len("quote"), lexer.pos)
	}
}
