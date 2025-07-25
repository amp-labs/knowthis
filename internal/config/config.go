package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port              string
	DatabaseURL       string
	SlackBotToken     string
	OpenAIAPIKey      string
	LogLevel          string
	LogFormat         string
	Environment       string
}

func Load() *Config {
	return &Config{
		Port:              os.Getenv("PORT"),
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		SlackBotToken:     os.Getenv("SLACK_BOT_TOKEN"),
		OpenAIAPIKey:      os.Getenv("OPENAI_API_KEY"),
		LogLevel:          os.Getenv("LOG_LEVEL"),
		LogFormat:         os.Getenv("LOG_FORMAT"),
		Environment:       os.Getenv("ENVIRONMENT"),
	}
}

func (c *Config) Validate() error {
	var errors []string

	if c.Port == "" {
		errors = append(errors, "PORT is required")
	}

	if c.SlackBotToken == "" {
		errors = append(errors, "SLACK_BOT_TOKEN is required")
	}


	if c.OpenAIAPIKey == "" {
		errors = append(errors, "OPENAI_API_KEY is required")
	}

	if c.DatabaseURL == "" {
		errors = append(errors, "DATABASE_URL is required")
	}

	if c.LogLevel == "" {
		errors = append(errors, "LOG_LEVEL is required")
	}

	if c.LogFormat == "" {
		errors = append(errors, "LOG_FORMAT is required")
	}

	if c.Environment == "" {
		errors = append(errors, "ENVIRONMENT is required")
	}

	// Optional validations
	if c.SlackBotToken != "" && !strings.HasPrefix(c.SlackBotToken, "xoxb-") {
		errors = append(errors, "SLACK_BOT_TOKEN must start with 'xoxb-'")
	}


	if c.LogLevel != "" {
		validLogLevels := []string{"DEBUG", "INFO", "WARN", "ERROR"}
		if !contains(validLogLevels, strings.ToUpper(c.LogLevel)) {
			errors = append(errors, "LOG_LEVEL must be one of: DEBUG, INFO, WARN, ERROR")
		}
	}

	if c.LogFormat != "" {
		validLogFormats := []string{"text", "json"}
		if !contains(validLogFormats, strings.ToLower(c.LogFormat)) {
			errors = append(errors, "LOG_FORMAT must be one of: text, json")
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("%s", errors[0])
	}

	return nil
}

func (c *Config) IsProduction() bool {
	return strings.ToLower(c.Environment) == "production"
}

func (c *Config) IsDevelopment() bool {
	return strings.ToLower(c.Environment) == "development"
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
