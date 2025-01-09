package main

import (
	"github.com/brendanryan/gmail-brita/internal/config"
)

// ExampleFilters returns a set of example Gmail filter configurations
func ExampleFilters() *config.Config {
	return &config.Config{
		Emails: []string{
			"me@example.com",
			"me@test.com",
		},
		Filters: []config.Filter{
			// Simple mailing list filter
			{
				Name: "Side Project List",
				Conditions: config.Conditions{
					Has: []string{"list:discuss@lists.some-side-project.org"},
				},
				Actions: config.Actions{
					Label:                 "some-side-project",
					ArchiveUnlessDirected: &config.ArchiveUnlessDirected{},
				},
			},
			// Complex filter chain for work emails
			{
				Name: "Important Robots",
				Conditions: config.Conditions{
					Has: []string{
						"list:robots@bigco.com",
						"subject:Important",
					},
				},
				Actions: config.Actions{
					Label: "work/robots/important",
				},
			},
			{
				Name: "Irrelevant Robots",
				Conditions: config.Conditions{
					Has: []string{
						"list:robots@bigco.com",
						"subject:Chunder",
					},
				},
				Actions: config.Actions{
					Label: "work/robots/irrelevant",
					ArchiveUnlessDirected: &config.ArchiveUnlessDirected{
						MarkRead: true,
					},
				},
			},
			{
				Name: "Meh Robots",
				Conditions: config.Conditions{
					Has: []string{
						"list:robots@bigco.com",
						"subject:Semirelevant",
					},
				},
				Actions: config.Actions{
					Label:                 "work/robots/meh",
					ArchiveUnlessDirected: &config.ArchiveUnlessDirected{},
				},
			},
			// Important personal emails
			{
				Name: "Family Emails",
				Conditions: config.Conditions{
					Has: []string{
						"from:mom@example.com",
						"from:dad@example.com",
						"from:sister@example.com",
					},
				},
				Actions: config.Actions{
					Label: "personal/family",
					Star:  true,
				},
			},
			// Financial notifications
			{
				Name: "Bank Statements",
				Conditions: config.Conditions{
					Has: []string{
						"from:statements@bank.com",
						"subject:\"Monthly Statement\"",
					},
				},
				Actions: config.Actions{
					Label:     "finance/bank-statements",
					Star:      true,
					NeverSpam: true,
				},
			},
			// Social media notifications
			{
				Name: "Social Updates",
				Conditions: config.Conditions{
					Has: []string{
						"from:notifications@twitter.com",
						"from:notification@linkedin.com",
						"from:notification@facebook.com",
					},
					HasNot: []string{
						"subject:\"Security alert\"",
						"subject:\"Login attempt\"",
					},
				},
				Actions: config.Actions{
					Label: "social/updates",
					ArchiveUnlessDirected: &config.ArchiveUnlessDirected{
						MarkRead: true,
					},
				},
			},
			// Shopping receipts and tracking
			{
				Name: "Shopping",
				Conditions: config.Conditions{
					Has: []string{
						"from:amazon.com",
						"from:orders@*.com",
						"subject:\"order confirmation\"",
						"subject:\"tracking number\"",
						"subject:\"shipped\"",
					},
				},
				Actions: config.Actions{
					Label: "shopping",
					Star:  true,
				},
			},
			// Calendar invites
			{
				Name: "Calendar",
				Conditions: config.Conditions{
					Has: []string{
						"filename:invite.ics",
						"subject:\"invited you to\"",
						"subject:\"calendar invitation\"",
					},
				},
				Actions: config.Actions{
					Label: "calendar",
					Star:  true,
				},
			},
			// Promotional emails
			{
				Name: "Promotions",
				Conditions: config.Conditions{
					Has: []string{
						"subject:\"% off\"",
						"subject:\"sale\"",
						"subject:\"deal\"",
						"subject:\"limited time\"",
					},
					HasNot: []string{
						"from:important-sender@example.com",
					},
				},
				Actions: config.Actions{
					Label: "promotions",
					ArchiveUnlessDirected: &config.ArchiveUnlessDirected{
						MarkRead: true,
					},
				},
			},
			// Travel itineraries
			{
				Name: "Travel",
				Conditions: config.Conditions{
					Has: []string{
						"subject:\"itinerary\"",
						"subject:\"booking confirmation\"",
						"subject:\"flight confirmation\"",
						"subject:\"hotel reservation\"",
						"from:*@airlines.com",
						"from:*@hotels.com",
					},
				},
				Actions: config.Actions{
					Label:     "travel",
					Star:      true,
					NeverSpam: true,
				},
			},
		},
	}
}
