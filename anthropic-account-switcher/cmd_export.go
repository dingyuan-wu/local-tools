package main

import (
	"fmt"
	"strings"
)

func runExport(cfg *Config, name string) error {
	if name == "" {
		return fmt.Errorf("account name is required")
	}
	acc, ok := cfg.Accounts[name]
	if !ok {
		return fmt.Errorf("account %q not found", name)
	}
	// Output to stdout for: eval "$(anthropic-account-switcher export prod)"
	fmt.Printf("export ANTHROPIC_AUTH_TOKEN=%s\n", quoteEnv(acc.AuthToken))
	if acc.OrgID != "" {
		fmt.Printf("export ANTHROPIC_ORG_ID=%s\n", quoteEnv(acc.OrgID))
	} else {
		fmt.Printf("unset ANTHROPIC_ORG_ID\n")
	}
	if acc.BaseURL != "" {
		fmt.Printf("export ANTHROPIC_BASE_URL=%s\n", quoteEnv(acc.BaseURL))
	} else {
		fmt.Printf("unset ANTHROPIC_BASE_URL\n")
	}
	// Prefer auth token mode when using this switcher.
	fmt.Printf("unset ANTHROPIC_API_KEY\n")
	fmt.Printf("export ANTHROPIC_ACCOUNT_SWITCHER_PROFILE=%s\n", quoteEnv(name))
	// Update current in config
	cfg.Current = name
	_ = saveConfig(cfg)
	return nil
}

func quoteEnv(s string) string {
	// Use single quotes; escape single quotes as '\''
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}

