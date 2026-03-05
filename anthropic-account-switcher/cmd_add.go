package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func runAdd(cfg *Config, name string) error {
	if name == "" {
		return fmt.Errorf("account name is required")
	}
	acc := Account{}
	// Prefer env if set
	if token := os.Getenv("ANTHROPIC_AUTH_TOKEN"); token != "" {
		acc.AuthToken = token
	}
	if org := os.Getenv("ANTHROPIC_ORG_ID"); org != "" {
		acc.OrgID = org
	}
	if baseURL := os.Getenv("ANTHROPIC_BASE_URL"); baseURL != "" {
		acc.BaseURL = baseURL
	}

	reader := bufio.NewReader(os.Stdin)
	if acc.AuthToken == "" {
		fmt.Fprint(os.Stderr, "ANTHROPIC_AUTH_TOKEN: ")
		fd := int(os.Stdin.Fd())
		raw, err := term.ReadPassword(fd)
		if err != nil {
			return fmt.Errorf("read auth token: %w", err)
		}
		acc.AuthToken = strings.TrimSpace(string(raw))
		fmt.Fprintln(os.Stderr)
	}
	if acc.OrgID == "" {
		fmt.Fprint(os.Stderr, "ANTHROPIC_ORG_ID (optional, press Enter to skip): ")
		line, _ := reader.ReadString('\n')
		acc.OrgID = strings.TrimSpace(line)
	}
	if acc.BaseURL == "" {
		fmt.Fprint(os.Stderr, "ANTHROPIC_BASE_URL (optional, press Enter to skip): ")
		line, _ := reader.ReadString('\n')
		acc.BaseURL = strings.TrimSpace(line)
	}

	cfg.Accounts[name] = acc
	cfg.Current = name
	if err := saveConfig(cfg); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Account %q added and set as current.\n", name)
	return nil
}

