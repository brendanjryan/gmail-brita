package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// LoadFromFile loads a filter configuration from a YAML file
func LoadFromFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &config, nil
}

// validateConfig checks that the configuration is valid
func validateConfig(config *Config) error {
	if len(config.Emails) == 0 {
		return fmt.Errorf("no email addresses specified")
	}

	if len(config.Filters) == 0 {
		return fmt.Errorf("no filters specified")
	}

	for i := 0; i < len(config.Filters); i++ {
		if err := validateFilter(&config.Filters[i], i); err != nil {
			return err
		}
	}

	return nil
}

// validateFilter checks that a filter configuration is valid
func validateFilter(filter *Filter, index int) error {
	if filter.Name == "" {
		return fmt.Errorf("filter %d has no name", index)
	}

	if len(filter.Conditions.Has) == 0 && len(filter.Conditions.HasNot) == 0 {
		return fmt.Errorf("filter %q has no conditions", filter.Name)
	}

	return nil
}
