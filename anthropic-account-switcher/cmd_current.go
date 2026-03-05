package main

import (
	"fmt"
	"os"
)

func runCurrent(cfg *Config) error {
	profile := os.Getenv("ANTHROPIC_ACCOUNT_SWITCHER_PROFILE")
	if profile != "" {
		fmt.Printf("Current profile: %s\n", profile)
		return nil
	}
	token := os.Getenv("ANTHROPIC_AUTH_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "Current profile: (none)")
		fmt.Fprintln(os.Stderr, "No ANTHROPIC_AUTH_TOKEN or ANTHROPIC_ACCOUNT_SWITCHER_PROFILE set.")
		return nil
	}
	// Match by auth token
	for name, acc := range cfg.Accounts {
		if acc.AuthToken == token {
			fmt.Printf("Current profile: %s\n", name)
			return nil
		}
	}
	// Unknown key: show prefix only
	prefix := token
	if len(prefix) > 12 {
		prefix = prefix[:12] + "..."
	}
	fmt.Printf("Current profile: unknown (ANTHROPIC_AUTH_TOKEN prefix: %s)\n", prefix)
	return nil
}

