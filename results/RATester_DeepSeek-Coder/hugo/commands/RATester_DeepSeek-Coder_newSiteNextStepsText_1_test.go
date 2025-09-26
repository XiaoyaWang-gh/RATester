package commands

import (
	"fmt"
	"testing"
)

func TestNewSiteNextStepsText_1(t *testing.T) {
	type args struct {
		path   string
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				path:   "/path/to/site",
				format: "md",
			},
			want: `Just a few more steps...

1. Change the current directory to /path/to/site.
2. Create or install a theme:
   - Create a new theme with the command "hugo new theme <THEMENAME>"
   - Or, install a theme from https://themes.gohugo.io/
3. Edit hugo.md, setting the "theme" property to the theme name.
4. Create new content with the command "hugo new content <SECTIONNAME>/<FILENAME>.md".
5. Start the embedded web server with the command "hugo server --buildDrafts".

See documentation at https://gohugo.io/.`,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &newCommand{}
			if got := c.newSiteNextStepsText(tt.args.path, tt.args.format); got != tt.want {
				t.Errorf("newSiteNextStepsText() = %v, want %v", got, tt.want)
			}
		})
	}
}
