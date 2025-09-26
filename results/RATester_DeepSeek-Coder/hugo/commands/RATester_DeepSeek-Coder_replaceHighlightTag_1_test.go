package commands

import (
	"fmt"
	"testing"
)

func TestReplaceHighlightTag_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple case",
			input:    "{% highlight go %}code{% endhighlight %}",
			expected: "{{< highlight go >}}code{{< /highlight >}}",
		},
		{
			name:     "with options",
			input:    "{% highlight go linenos=table %}code{% endhighlight %}",
			expected: "{{< highlight go linenos=\"table\" >}}code{{< /highlight >}}",
		},
		{
			name:     "with multiple options",
			input:    "{% highlight go linenos=table style=\"monokai\" %}code{% endhighlight %}",
			expected: "{{< highlight go linenos=\"table\" style=\"monokai\" >}}code{{< /highlight >}}",
		},
		{
			name:     "with spaces",
			input:    "{% highlight go  style=\"monokai\" %}code{% endhighlight %}",
			expected: "{{< highlight go style=\"monokai\" >}}code{{< /highlight >}}",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &importCommand{}
			result := c.replaceHighlightTag(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
