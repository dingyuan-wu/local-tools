package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	configDir  = "~.anthropic-account-switcher"
	configFile = "accounts.json"
)

// Account holds Anthropic credentials for one account.
type Account struct {
	AuthToken string `json:"anthropic_auth_token"`
	OrgID     string `json:"anthropic_org_id,omitempty"`
	BaseURL   string `json:"anthropic_base_url,omitempty"`
}

// Config is the on-disk store.
type Config struct {
	Accounts map[string]Account `json:"accounts"`
	Current  string             `json:"current,omitempty"`
}

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home dir: %w", err)
	}
	dir := configDir
	if strings.HasPrefix(dir, "~") {
		dir = strings.TrimPrefix(dir, "~")
	}
	return filepath.Join(home, dir, configFile), nil
}

func ensureConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home dir: %w", err)
	}
	dirPart := configDir
	if strings.HasPrefix(dirPart, "~") {
		dirPart = strings.TrimPrefix(dirPart, "~")
	}
	dir := filepath.Join(home, dirPart)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", fmt.Errorf("create config dir: %w", err)
	}
	return dir, nil
}

func loadConfig() (*Config, string, error) {
	p, err := configPath()
	if err != nil {
		return nil, "", err
	}
	data, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{Accounts: make(map[string]Account)}, p, nil
		}
		return nil, "", fmt.Errorf("read config: %w", err)
	}
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, "", fmt.Errorf("parse config: %w", err)
	}
	if c.Accounts == nil {
		c.Accounts = make(map[string]Account)
	}
	return &c, p, nil
}

func saveConfig(c *Config) error {
	_, err := ensureConfigDir()
	if err != nil {
		return err
	}
	p, err := configPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}
	if err := os.WriteFile(p, data, 0o600); err != nil {
		return fmt.Errorf("write config: %w", err)
	}
	return nil
}

