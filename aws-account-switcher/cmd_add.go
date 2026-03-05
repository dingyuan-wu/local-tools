package switcher

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
	if id := os.Getenv("AWS_ACCESS_KEY_ID"); id != "" {
		acc.AccessKeyID = id
	}
	if secret := os.Getenv("AWS_SECRET_ACCESS_KEY"); secret != "" {
		acc.SecretAccessKey = secret
	}
	if token := os.Getenv("AWS_SESSION_TOKEN"); token != "" {
		acc.SessionToken = token
	}

	reader := bufio.NewReader(os.Stdin)
	if acc.AccessKeyID == "" {
		fmt.Fprint(os.Stderr, "AWS_ACCESS_KEY_ID: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("read access key id: %w", err)
		}
		acc.AccessKeyID = strings.TrimSpace(line)
	}
	if acc.SecretAccessKey == "" {
		fmt.Fprint(os.Stderr, "AWS_SECRET_ACCESS_KEY: ")
		fd := int(os.Stdin.Fd())
		raw, err := term.ReadPassword(fd)
		if err != nil {
			return fmt.Errorf("read secret key: %w", err)
		}
		acc.SecretAccessKey = string(raw)
		fmt.Fprintln(os.Stderr)
	}
	if acc.SessionToken == "" {
		fmt.Fprint(os.Stderr, "AWS_SESSION_TOKEN (optional, press Enter to skip): ")
		line, _ := reader.ReadString('\n')
		acc.SessionToken = strings.TrimSpace(line)
	}

	cfg.Accounts[name] = acc
	cfg.Current = name
	if err := saveConfig(cfg); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Account %q added and set as current.\n", name)
	return nil
}
