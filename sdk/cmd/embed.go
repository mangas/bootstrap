package cmd

import (
	"github.com/getcouragenow/bootstrap/sdk/pkg/statics"
	"github.com/spf13/cobra"
)

func NewInitBoilerplateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Args: cobra.ExactArgs(1),
		Short: "Write boilerplates to your current directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			bp, err := statics.NewRootBoilerplate()
			if err != nil {
				return err
			}
			if err = bp.WriteAllFiles(args[0]); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}