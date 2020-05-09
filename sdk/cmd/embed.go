package cmd

import (
	"github.com/getcouragenow/bootstrap/sdk/pkg/statics"
	"github.com/spf13/cobra"
)

func NewInitBoilerplateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [core|lyft|tool|all] <output_dir>",
		Args: cobra.ExactArgs(2),
		Short: "Write boilerplates to your current directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			bp, err := statics.NewBPAsset(args[0])
			if err != nil {
				return err
			}
			if err = bp.WriteAllFiles(args[1]); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}