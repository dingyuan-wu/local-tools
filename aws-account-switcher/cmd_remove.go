package switcher

import (
	"fmt"
	"os"
)

func runRemove(cfg *Config, name string) error {
	if name == "" {
		return fmt.Errorf("account name is required")
	}
	if _, ok := cfg.Accounts[name]; !ok {
		return fmt.Errorf("account %q not found", name)
	}
	delete(cfg.Accounts, name)
	if cfg.Current == name {
		cfg.Current = ""
	}
	if err := saveConfig(cfg); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Account %q removed.\n", name)
	return nil
}
