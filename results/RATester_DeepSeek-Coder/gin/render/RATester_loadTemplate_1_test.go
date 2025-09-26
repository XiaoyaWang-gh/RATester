package render

// import (
// 	"fmt"
// 	"html/template"
// 	"reflect"
// 	"testing"
// )

// func TestLoadTemplate_1(t *testing.T) {
// 	testCases := []struct {
// 		name     string
// 		r        HTMLDebug
// 		expected *template.Template
// 	}{
// 		{
// 			name: "Test with files",
// 			r: HTMLDebug{
// 				Files: []string{"file1.html", "file2.html"},
// 			},
// 			expected: template.Must(template.New("").ParseFiles("file1.html", "file2.html")),
// 		},
// 		{
// 			name: "Test with glob pattern",
// 			r: HTMLDebug{
// 				Glob: "*.html",
// 			},
// 			expected: template.Must(template.New("").ParseGlob("*.html")),
// 		},
// 		{
// 			name: "Test without files or glob pattern",
// 			r:    HTMLDebug{},
// 			expected: func() *template.Template {
// 				defer func() {
// 					if r := recover(); r == nil {
// 						t.Errorf("The code did not panic")
// 					}
// 				}()
// 				return nil
// 			}(),
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {

// 			defer func() {
// 				if r := recover(); r != nil {
// 					fmt.Println("Recovered in main", r)
// 				}
// 			}()

// 			result := tc.r.loadTemplate()
// 			if !reflect.DeepEqual(result, tc.expected) {
// 				t.Errorf("Expected %v, got %v", tc.expected, result)
// 			}
// 		})
// 	}
// }
