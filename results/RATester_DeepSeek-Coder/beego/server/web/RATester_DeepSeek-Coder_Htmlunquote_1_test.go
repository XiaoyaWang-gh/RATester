package web

import (
	"fmt"
	"testing"
)

func TestHtmlunquote_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{input: "&lt;p&gt;This is a paragraph.&lt;/p&gt;", expected: "<p>This is a paragraph.</p>"},
		{input: "&amp;lt;p&amp;gt;This is another paragraph.&amp;lt;/p&amp;gt;", expected: "<p>This is another paragraph.</p>"},
		{input: "&lt;p&gt; &lt;/p&gt;", expected: "<p> </p>"},
		{input: "&lt;p&gt; &amp;nbsp; &lt;/p&gt;", expected: "<p> &nbsp; </p>"},
		{input: "&lt;p&gt; &amp;#160; &lt;/p&gt;", expected: "<p> &#160; </p>"},
	}

	for _, test := range tests {
		result := Htmlunquote(test.input)
		if result != test.expected {
			t.Errorf("Htmlunquote(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
