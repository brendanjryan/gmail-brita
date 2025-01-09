package filter

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// Set represents a collection of Gmail filters
type Set struct {
	Emails  []string
	Filters []*Filter
}

// Filter represents a Gmail filter
type Filter struct {
	HasWords         []string
	DoesNotHaveWords []string
	Labels           []string
	Archive          bool
	MarkRead         bool
	Star             bool
	NeverSpam        bool
}

// NewFilterSet creates a new filter set with the given email addresses
func NewFilterSet(emails []string) *Set {
	return &Set{
		Emails:  emails,
		Filters: make([]*Filter, 0),
	}
}

// AddFilter adds a new filter to the set
func (s *Set) AddFilter() *Filter {
	filter := &Filter{
		HasWords:         make([]string, 0),
		DoesNotHaveWords: make([]string, 0),
		Labels:           make([]string, 0),
	}
	s.Filters = append(s.Filters, filter)
	return filter
}

// ToXML converts the filter set to Gmail's XML format
func (s *Set) ToXML() ([]byte, error) {
	feed := &Feed{
		XMLName:  xml.Name{Space: "http://www.w3.org/2005/Atom", Local: "feed"},
		XMLNS:    "http://www.w3.org/2005/Atom",
		XMLNSApp: "http://schemas.google.com/apps/2006",
		Title:    "Mail Filters",
		ID:       fmt.Sprintf("tag:mail.google.com,2008:filters:%s", s.Emails[0]),
		Updated:  time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		Author: Author{
			Name:  s.Emails[0],
			Email: s.Emails[0],
		},
		Entries: make([]Entry, 0),
	}

	for i, filter := range s.Filters {
		entry := Entry{
			Category: Category{Term: "filter"},
			Title:    "Mail Filter",
			ID:       fmt.Sprintf("tag:mail.google.com,2008:filter:%d", i+1),
			Updated:  feed.Updated,
			Content:  "",
		}

		if len(filter.HasWords) > 0 {
			entry.Properties = append(entry.Properties, Property{
				Name:  "hasTheWord",
				Value: strings.Join(filter.HasWords, " AND "),
			})
		}

		if len(filter.DoesNotHaveWords) > 0 {
			entry.Properties = append(entry.Properties, Property{
				Name:  "doesNotHaveWord",
				Value: strings.Join(filter.DoesNotHaveWords, " OR "),
			})
		}

		if len(filter.Labels) > 0 {
			for _, label := range filter.Labels {
				entry.Properties = append(entry.Properties, Property{
					Name:  "label",
					Value: label,
				})
			}
		}

		if filter.Archive {
			entry.Properties = append(entry.Properties, Property{
				Name:  "shouldArchive",
				Value: "true",
			})
		}

		if filter.MarkRead {
			entry.Properties = append(entry.Properties, Property{
				Name:  "shouldMarkAsRead",
				Value: "true",
			})
		}

		if filter.Star {
			entry.Properties = append(entry.Properties, Property{
				Name:  "shouldStar",
				Value: "true",
			})
		}

		if filter.NeverSpam {
			entry.Properties = append(entry.Properties, Property{
				Name:  "neverSpam",
				Value: "true",
			})
		}

		feed.Entries = append(feed.Entries, entry)
	}

	return xml.MarshalIndent(feed, "", "  ")
}

// Feed represents the root element of Gmail's filter XML
type Feed struct {
	XMLName  xml.Name `xml:"feed"`
	XMLNS    string   `xml:"xmlns,attr"`
	XMLNSApp string   `xml:"xmlns:apps,attr"`
	Title    string   `xml:"title"`
	ID       string   `xml:"id"`
	Updated  string   `xml:"updated"`
	Author   Author   `xml:"author"`
	Entries  []Entry  `xml:"entry"`
}

// Author represents the author element in Gmail's filter XML
type Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

// Entry represents a filter entry in Gmail's filter XML
type Entry struct {
	Category   Category   `xml:"category"`
	Title      string     `xml:"title"`
	ID         string     `xml:"id"`
	Updated    string     `xml:"updated"`
	Content    string     `xml:"content"`
	Properties []Property `xml:"apps:property"`
}

// Category represents a category element in Gmail's filter XML
type Category struct {
	Term string `xml:"term,attr"`
}

// Property represents a property element in Gmail's filter XML
type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
