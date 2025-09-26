package provider

import (
	"context"
	"fmt"

	"github.com/shawn-hurley/jsonrpc-golang/engine"
	"github.com/shawn-hurley/jsonrpc-golang/provider/golang"
	"github.com/shawn-hurley/jsonrpc-golang/provider/lib"
)

// For some period of time during POC this will be in tree, in the future we need to write something that can do this w/ external binaries
type Client interface {
	Capabilities() ([]string, error)

	// Block until initialized
	Init(ctx context.Context, workspace string) error

	GetAllCompletions(row int, column int, file_path string) string

	GetDefinition(row int, column int, file_path string) string

	Evaluate(cap string, conditionInfo interface{}) (lib.ProviderEvaluateResponse, error)
}

type ProviderCondition struct {
	Client
	Capability    string
	ConditionInfo interface{}
}

func (p *ProviderCondition) Evaluate() (engine.CondtionResponse, error) {
	resp, err := p.Client.Evaluate(p.Capability, p.ConditionInfo)
	if err != nil {
		// If an error always just return the empty
		return engine.CondtionResponse{}, err
	}

	return engine.CondtionResponse{
		Passed:              resp.Passed,
		ConditionHitContext: resp.ConditionHitContext,
	}, nil

}

func (p *ProviderCondition) GetAllCompletions(row int, column int, file_path string) string {
	res := p.Client.GetAllCompletions(row, column, file_path)
	return res
}

func (p *ProviderCondition) GetDefinition(row int, column int, file_path string) string {
	res := p.Client.GetDefinition(row, column, file_path)
	return res
}

// We need some wrapper that can deal with out of tree providers, this will be a call, that will mock it out, but go against in tree.
func GetProviderClient(config lib.Config) (Client, error) {
	switch config.Name {
	case "go":
		return golang.NewGolangProvider(config), nil
	// case "java":
	// 	return java.NewJavaProvider(config), nil
	default:
		return nil, fmt.Errorf("unknown and invalid provider client")
	}
}
