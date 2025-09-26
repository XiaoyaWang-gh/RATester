package commands

import (
	"fmt"
	"testing"
)

func TestreplaceHighlightTag_1(t *testing.T) {
	type args struct {
		match string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				match: "{% highlight go %}",
			},
			want: "{{< highlight go >}}",
		},
		{
			name: "Test case 2",
			args: args{
				match: "{% highlight go linenos=table %}",
			},
			want: "{{< highlight go linenos=\"table\" >}}",
		},
		{
			name: "Test case 3",
			args: args{
				match: "{% highlight go linenos=table,linenostart=10 %}",
			},
			want: "{{< highlight go linenos=\"table,linenostart=10\" >}}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &importCommand{}
			if got := c.replaceHighlightTag(tt.args.match); got != tt.want {
				t.Errorf("replaceHighlightTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
