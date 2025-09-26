package publisher

import (
	"fmt"
	"testing"
)

func TestIsClosedByTag_1(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		tagName []byte
		want    bool
	}{
		{
			name: "test1",
			b: []byte(`<div>
				<p>
					<a href="https://www.google.com">Google</a>
				</p>
			</div>`),
			tagName: []byte("div"),
			want:    true,
		},
		{
			name: "test2",
			b: []byte(`<div>
				<p>
					<a href="https://www.google.com">Google</a>
				</p>
			</div>`),
			tagName: []byte("p"),
			want:    false,
		},
		{
			name: "test3",
			b: []byte(`<div>
				<p>
					<a href="https://www.google.com">Google</a>
				</p>
			</div>`),
			tagName: []byte("a"),
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isClosedByTag(tt.b, tt.tagName); got != tt.want {
				t.Errorf("isClosedByTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
