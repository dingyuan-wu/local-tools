package switcher

import (
	"fmt"
	"os"
	"sort"
)

func runList(cfg *Config) error {
	if len(cfg.Accounts) == 0 {
		fmt.Fprintln(os.Stderr, "No accounts saved. Use 'add <name>' to add one.")
		return nil
	}
	names := make([]string, 0, len(cfg.Accounts))
	for n := range cfg.Accounts {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		marker := " "
		if n == cfg.Current {
			marker = "*"
		}
		fmt.Printf("%s %s\n", marker, n)
	}
	return nil
}
