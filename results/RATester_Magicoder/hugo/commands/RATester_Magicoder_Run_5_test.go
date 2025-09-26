package commands

import (
	"context"
	"fmt"
	"testing"

	"github.com/bep/simplecobra"
)

func TestRun_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	cd := &simplecobra.Commandeer{
		Command: &serverCommand{
			r: &rootCommand{
				Printf: func(format string, v ...interface{}) {
					t.Logf(format, v...)
				},
			},
		},
	}
	args := []string{}
	err := cd.Command.Run(ctx, cd, args)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
