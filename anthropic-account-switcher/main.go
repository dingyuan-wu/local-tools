package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "anthropic-account-switcher",
	Short: "Manage and switch between multiple Anthropic API accounts",
	Long: `Manage local Anthropic API keys for multiple accounts,
switch with one command, and query the current account.`,
}

var addCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add or update an account (interactive or from env)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, _, err := loadConfig()
		if err != nil {
			return err
		}
		return runAdd(cfg, args[0])
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved account names",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, _, err := loadConfig()
		if err != nil {
			return err
		}
		return runList(cfg)
	},
}

var useCmd = &cobra.Command{
	Use:   "use [name]",
	Short: "Switch to an account (must use with eval to apply in current shell)",
	Long:  `Outputs export statements to stdout. To apply in your current shell, run: eval "$(anthropic-account-switcher use [name])"`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, _, err := loadConfig()
		if err != nil {
			return err
		}
		name := args[0]
		if _, ok := cfg.Accounts[name]; !ok {
			return fmt.Errorf("account %q not found", name)
		}
		fmt.Fprintf(os.Stderr, "To apply in this shell, run:\n  eval \"$(anthropic-account-switcher use %s)\"\n", name)
		return runExport(cfg, name)
	},
}

var exportCmd = &cobra.Command{
	Use:   "export [name]",
	Short: "Output shell export statements for eval/source",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, _, err := loadConfig()
		if err != nil {
			return err
		}
		return runExport(cfg, args[0])
	},
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show which account is currently in use",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, _, err := loadConfig()
		if err != nil {
			return err
		}
		return runCurrent(cfg)
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a saved account",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, _, err := loadConfig()
		if err != nil {
			return err
		}
		return runRemove(cfg, args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd, listCmd, useCmd, exportCmd, currentCmd, removeCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

