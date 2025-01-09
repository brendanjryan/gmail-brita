package config

// Config represents the top-level YAML configuration
type Config struct {
	Emails  []string `yaml:"emails"`
	Filters []Filter `yaml:"filters"`
}

// Filter represents a single Gmail filter configuration
type Filter struct {
	Name       string     `yaml:"name"`
	Conditions Conditions `yaml:"conditions"`
	Actions    Actions    `yaml:"actions"`
}

// Conditions represents the conditions for a filter
type Conditions struct {
	Has    []string `yaml:"has,omitempty"`
	HasNot []string `yaml:"has_not,omitempty"`
}

// Actions represents the actions for a filter
type Actions struct {
	Label                 string                 `yaml:"label,omitempty"`
	Archive               bool                   `yaml:"archive,omitempty"`
	MarkRead              bool                   `yaml:"mark_read,omitempty"`
	Star                  bool                   `yaml:"star,omitempty"`
	NeverSpam             bool                   `yaml:"never_spam,omitempty"`
	ArchiveUnlessDirected *ArchiveUnlessDirected `yaml:"archive_unless_directed,omitempty"`
}

// ArchiveUnlessDirected represents the archive_unless_directed action parameters
type ArchiveUnlessDirected struct {
	MarkRead bool `yaml:"mark_read,omitempty"`
}
