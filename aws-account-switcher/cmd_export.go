package switcher

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
	// Output to stdout for: eval "$(aws-account-switcher export prod)"
	fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", quoteEnv(acc.AccessKeyID))
	fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", quoteEnv(acc.SecretAccessKey))
	if acc.SessionToken != "" {
		fmt.Printf("export AWS_SESSION_TOKEN=%s\n", quoteEnv(acc.SessionToken))
	} else {
		fmt.Printf("unset AWS_SESSION_TOKEN\n")
	}
	fmt.Printf("export AWS_ACCOUNT_SWITCHER_PROFILE=%s\n", quoteEnv(name))
	// Update current in config
	cfg.Current = name
	_ = saveConfig(cfg)
	return nil
}

func quoteEnv(s string) string {
	// Use single quotes; escape single quotes as '\''
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}
