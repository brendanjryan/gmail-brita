package filter

import (
	"fmt"
)

// Builder provides a fluent interface for building Gmail filters
type Builder struct {
	filter *Filter
	set    *Set
	chain  []*Filter
}

// NewBuilder creates a new filter builder
func NewBuilder(set *Set) *Builder {
	return &Builder{
		filter: set.AddFilter(),
		set:    set,
		chain:  make([]*Filter, 0),
	}
}

// Has adds positive match conditions to the filter
func (b *Builder) Has(words []string) *Builder {
	b.filter.HasWords = append(b.filter.HasWords, words...)
	return b
}

// HasNot adds negative match conditions to the filter
func (b *Builder) HasNot(words []string) *Builder {
	b.filter.DoesNotHaveWords = append(b.filter.DoesNotHaveWords, words...)
	return b
}

// Label adds a label action to the filter
func (b *Builder) Label(label string) *Builder {
	b.filter.Labels = append(b.filter.Labels, label)
	return b
}

// Archive adds an archive action to the filter
func (b *Builder) Archive() *Builder {
	b.filter.Archive = true
	return b
}

// MarkRead adds a mark-as-read action to the filter
func (b *Builder) MarkRead() *Builder {
	b.filter.MarkRead = true
	return b
}

// Star adds a star action to the filter
func (b *Builder) Star() *Builder {
	b.filter.Star = true
	return b
}

// NeverSpam adds a never-mark-as-spam action to the filter
func (b *Builder) NeverSpam() *Builder {
	b.filter.NeverSpam = true
	return b
}

// ArchiveUnlessDirectedOption represents an option for the ArchiveUnlessDirected method
type ArchiveUnlessDirectedOption func(*Filter)

// WithMarkRead returns an option to mark messages as read when archiving
func WithMarkRead(markRead bool) ArchiveUnlessDirectedOption {
	return func(f *Filter) {
		f.MarkRead = markRead
	}
}

// ArchiveUnlessDirected creates a chain filter that archives messages unless they are directed to the user
func (b *Builder) ArchiveUnlessDirected(opts ...ArchiveUnlessDirectedOption) *Builder {
	// Create a new filter for archiving
	archiveFilter := b.set.AddFilter()
	archiveFilter.HasWords = b.filter.HasWords
	archiveFilter.Archive = true

	// Add "to:" and "cc:" exclusions for each email
	for _, email := range b.set.Emails {
		archiveFilter.DoesNotHaveWords = append(
			archiveFilter.DoesNotHaveWords,
			fmt.Sprintf("to:%s", email),
			fmt.Sprintf("cc:%s", email),
		)
	}

	// Apply options
	for _, opt := range opts {
		opt(archiveFilter)
	}

	// Add to chain
	b.chain = append(b.chain, archiveFilter)
	return b
}

// Otherwise starts a new filter chain branch
func (b *Builder) Otherwise() *Builder {
	// Create inverse conditions from the previous filter
	notConditions := make([]string, 0)
	for _, has := range b.filter.HasWords {
		notConditions = append(notConditions, fmt.Sprintf("-%s", has))
	}

	// Create a new filter with the inverse conditions
	newFilter := b.set.AddFilter()
	newFilter.HasWords = notConditions

	return &Builder{
		filter: newFilter,
		set:    b.set,
		chain:  b.chain,
	}
}
