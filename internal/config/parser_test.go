package config

import (
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		wantErr bool
	}{
		{
			name:    "simple config",
			file:    "../testdata/filters/simple.yaml",
			wantErr: false,
		},
		{
			name:    "complex config",
			file:    "../testdata/filters/complex.yaml",
			wantErr: false,
		},
		{
			name:    "invalid config",
			file:    "../testdata/filters/invalid.yaml",
			wantErr: true,
		},
		{
			name:    "nonexistent file",
			file:    "nonexistent.yaml",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := LoadFromFile(tt.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && cfg == nil {
				t.Error("LoadFromFile() returned nil config without error")
			}
		})
	}
}

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: &Config{
				Emails: []string{"me@example.com"},
				Filters: []Filter{
					{
						Name: "Test Filter",
						Conditions: Conditions{
							Has: []string{"test"},
						},
						Actions: Actions{
							Label: "test",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "missing emails",
			config: &Config{
				Emails: []string{},
				Filters: []Filter{
					{
						Name: "Test Filter",
						Conditions: Conditions{
							Has: []string{"test"},
						},
						Actions: Actions{
							Label: "test",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing filters",
			config: &Config{
				Emails:  []string{"me@example.com"},
				Filters: []Filter{},
			},
			wantErr: true,
		},
		{
			name: "filter without name",
			config: &Config{
				Emails: []string{"me@example.com"},
				Filters: []Filter{
					{
						Conditions: Conditions{
							Has: []string{"test"},
						},
						Actions: Actions{
							Label: "test",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "filter without conditions",
			config: &Config{
				Emails: []string{"me@example.com"},
				Filters: []Filter{
					{
						Name: "Test Filter",
						Actions: Actions{
							Label: "test",
						},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateConfig(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
