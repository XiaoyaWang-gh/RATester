package tcp

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/rules"
)

func TestAddRule_1(t *testing.T) {
	type args struct {
		rule  *rules.Tree
		funcs matcherFuncs
	}
	tests := []struct {
		name    string
		m       *matchersTree
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &matchersTree{}
			if err := m.addRule(tt.args.rule, tt.args.funcs); (err != nil) != tt.wantErr {
				t.Errorf("matchersTree.addRule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
