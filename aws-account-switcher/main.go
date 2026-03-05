package switcher

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aws-account-switcher",
	Short: "Manage and switch between multiple AWS account credentials",
	Long: `Manage local AWS Access Key (AK) environment variables for multiple accounts,
switch with one command, and query the current account.`,
}

var addCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add or update an account (interactive or from env)",
	Args:  cobra.ExactArgs(1),
	RunE:  runAddCmd,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved account names",
	RunE:  runListCmd,
}

var useCmd = &cobra.Command{
	Use:   "use [name]",
	Short: "Switch to an account (must use with eval to apply in current shell)",
	Long:  `Outputs export statements to stdout. To apply in your current shell, run: eval "$(aws-account-switcher use [name])"`,
	Args:  cobra.ExactArgs(1),
	RunE:  runUseCmd,
}

var exportCmd = &cobra.Command{
	Use:   "export [name]",
	Short: "Output shell export statements for eval/source",
	Args:  cobra.ExactArgs(1),
	RunE:  runExportCmd,
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show which account is currently in use",
	RunE:  runCurrentCmd,
}

var removeCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a saved account",
	Args:  cobra.ExactArgs(1),
	RunE:  runRemoveCmd,
}

func runAddCmd(cmd *cobra.Command, args []string) error {
	cfg, _, err := loadConfig()
	if err != nil {
		return err
	}
	return runAdd(cfg, args[0])
}

func runListCmd(cmd *cobra.Command, args []string) error {
	cfg, _, err := loadConfig()
	if err != nil {
		return err
	}
	return runList(cfg)
}

func runUseCmd(cmd *cobra.Command, args []string) error {
	cfg, _, err := loadConfig()
	if err != nil {
		return err
	}
	name := args[0]
	if _, ok := cfg.Accounts[name]; !ok {
		return fmt.Errorf("account %q not found", name)
	}
	fmt.Fprintf(os.Stderr, "To apply in this shell, run:\n  eval \"$(aws-account-switcher use %s)\"\n", name)
	return runExport(cfg, name)
}

func runExportCmd(cmd *cobra.Command, args []string) error {
	cfg, _, err := loadConfig()
	if err != nil {
		return err
	}
	return runExport(cfg, args[0])
}

func runCurrentCmd(cmd *cobra.Command, args []string) error {
	cfg, _, err := loadConfig()
	if err != nil {
		return err
	}
	return runCurrent(cfg)
}

func runRemoveCmd(cmd *cobra.Command, args []string) error {
	cfg, _, err := loadConfig()
	if err != nil {
		return err
	}
	return runRemove(cfg, args[0])
}

func init() {
	rootCmd.AddCommand(addCmd, listCmd, useCmd, exportCmd, currentCmd, removeCmd)
}

// Execute runs the root command. It is intended to be called by the cmd/main package.
func Execute() error {
	return rootCmd.Execute()
}
