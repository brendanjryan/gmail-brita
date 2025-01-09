package britta

import (
	"github.com/brendanryan/gmail-brita/internal/config"
	"github.com/brendanryan/gmail-brita/internal/filter"
)

// GenerateXML generates Gmail filter XML from a configuration
func GenerateXML(cfg *config.Config) ([]byte, error) {
	// Create filter set
	set := filter.NewFilterSet(cfg.Emails)

	// Build filters
	for _, f := range cfg.Filters {
		builder := filter.NewBuilder(set)

		// Add conditions
		if len(f.Conditions.Has) > 0 {
			builder.Has(f.Conditions.Has)
		}
		if len(f.Conditions.HasNot) > 0 {
			builder.HasNot(f.Conditions.HasNot)
		}

		// Add actions
		if f.Actions.Label != "" {
			builder.Label(f.Actions.Label)
		}
		if f.Actions.Archive {
			builder.Archive()
		}
		if f.Actions.MarkRead {
			builder.MarkRead()
		}
		if f.Actions.Star {
			builder.Star()
		}
		if f.Actions.NeverSpam {
			builder.NeverSpam()
		}
		if f.Actions.ArchiveUnlessDirected != nil {
			var opts []filter.ArchiveUnlessDirectedOption
			if f.Actions.ArchiveUnlessDirected.MarkRead {
				opts = append(opts, filter.WithMarkRead(true))
			}
			builder.ArchiveUnlessDirected(opts...)
		}
	}

	// Generate XML
	return set.ToXML()
}
