package main

import (
	"context"
	"encoding/json"
	"flag"
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
	listenAddr := flag.String("listen", "127.0.0.1:2000", "listen address for the jsonrpc gopls service")
	flag.Parse()

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

	fmt.Printf("Server started at %s\n", *listenAddr)
	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		fmt.Printf("\n%v\n", err)
		os.Exit(1)
	}
}
