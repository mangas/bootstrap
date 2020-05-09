package cmd

import (
	"github.com/spf13/cobra"
)

const (
	short = `install specified tool to User's PATH, see 'list' command for list of tools`
)

func NewInstallToolsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "install <tool_name>",
		Short: short,
		Long: short,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
