package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/shawn-hurley/jsonrpc-golang/parser"
	"github.com/shawn-hurley/jsonrpc-golang/provider"
	"github.com/shawn-hurley/jsonrpc-golang/provider/lib"
)

const (
	SETTING_FILE_PATH = "./provider_settings.json"
	RULES_FILE_PATH   = "./rule-example.json"
)

var (
	providers map[string]provider.Client
)

type Request struct {
	Row       int    `json:"row"`
	Column    int    `json:"column"`
	Workspace string `json:"workspace"`
	Path      string `json:"path"`
}

func main() {
	ctx := context.Background()

	// Get the configs
	configs, err := lib.GetConfig(SETTING_FILE_PATH)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		os.Exit(1)
	}

	providers = map[string]provider.Client{}

	for _, config := range configs {
		provider, err := provider.GetProviderClient(config)
		if err != nil {
			fmt.Printf("\n%v\n", err)
			os.Exit(1)
		}
		providers[config.Name] = provider
	}

	parser := parser.RuleParser{
		ProviderNameToClient: providers,
	}

	_, needProviders, err := parser.LoadRules(RULES_FILE_PATH)

	if err != nil {
		fmt.Printf("\n%v\n", err)
		os.Exit(1)
	}
	provider := needProviders[0]

	http.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		fmt.Println("init")
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = provider.Init(ctx, req.Workspace)

		if err != nil {
			fmt.Printf("\n%v\n", err)
			os.Exit(1)
		}
	})

	http.HandleFunc("/completions", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := provider.GetAllCompletions(req.Row, req.Column, req.Path)
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/definition", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(req)
		res := provider.GetDefinition(req.Row, req.Column, req.Path)
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/publishDiagnostics", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := provider.GetAllCompletions(req.Row, req.Column, req.Path)
		json.NewEncoder(w).Encode(res)
	})

	fmt.Println("Server started at :2000")
	http.ListenAndServe(":2000", nil)
}
