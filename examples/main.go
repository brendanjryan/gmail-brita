package main

import (
	"fmt"
	"os"

	"github.com/brendanryan/gmail-brita/internal/config"
	"github.com/brendanryan/gmail-brita/pkg/britta"
)

func main() {
	// Create configuration
	cfg := &config.Config{
		Emails: []string{"me@example.com"},
		Filters: []config.Filter{
			{
				Name: "Example Filter",
				Conditions: config.Conditions{
					Has: []string{"list:example@list.com"},
				},
				Actions: config.Actions{
					Label: "example-list",
					ArchiveUnlessDirected: &config.ArchiveUnlessDirected{
						MarkRead: true,
					},
				},
			},
		},
	}

	// Generate XML
	xml, err := britta.GenerateXML(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating XML: %v\n", err)
		os.Exit(1)
	}

	// Write to stdout
	fmt.Println(string(xml))
}
