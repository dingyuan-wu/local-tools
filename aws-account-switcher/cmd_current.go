package switcher

import (
	"fmt"
	"os"
)

func runCurrent(cfg *Config) error {
	profile := os.Getenv("AWS_ACCOUNT_SWITCHER_PROFILE")
	if profile != "" {
		fmt.Printf("Current profile: %s\n", profile)
		return nil
	}
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	if accessKey == "" {
		fmt.Fprintln(os.Stderr, "Current profile: (none)")
		fmt.Fprintln(os.Stderr, "No AWS_ACCESS_KEY_ID or AWS_ACCOUNT_SWITCHER_PROFILE set.")
		return nil
	}
	// Match by access key id
	for name, acc := range cfg.Accounts {
		if acc.AccessKeyID == accessKey {
			fmt.Printf("Current profile: %s\n", name)
			return nil
		}
	}
	// Unknown key: show prefix only
	prefix := accessKey
	if len(prefix) > 12 {
		prefix = prefix[:12] + "..."
	}
	fmt.Printf("Current profile: unknown (AWS_ACCESS_KEY_ID prefix: %s)\n", prefix)
	return nil
}
