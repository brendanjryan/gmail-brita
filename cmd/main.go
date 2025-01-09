package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/brendanryan/gmail-brita/internal/config"
	"github.com/brendanryan/gmail-brita/pkg/britta"
)

func main() {
	var (
		configFile string
		outputFile string
	)

	flag.StringVar(&configFile, "config", "", "Path to YAML config file")
	flag.StringVar(&outputFile, "out", "", "Path to output XML file")
	flag.Parse()

	if configFile == "" {
		fmt.Fprintln(os.Stderr, "Error: config file is required")
		flag.Usage()
		os.Exit(1)
	}

	if outputFile == "" {
		fmt.Fprintln(os.Stderr, "Error: output file is required")
		flag.Usage()
		os.Exit(1)
	}

	// Load configuration
	cfg, err := config.LoadFromFile(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Generate XML
	xml, err := britta.GenerateXML(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating XML: %v\n", err)
		os.Exit(1)
	}

	// Write output
	if err := os.WriteFile(outputFile, xml, 0600); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}
}
