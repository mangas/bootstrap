package cmd

import "github.com/spf13/cobra"

func NewListToolsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list -[a|i]",
		Short: "List all tools available or currently installed",
		Long: "List all tools available (-a), or (-i) for list installed tools",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
