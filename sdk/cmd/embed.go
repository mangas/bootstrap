package cmd

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/spf13/cobra"
)

func newInitBoilerplateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Write boilerplates to your current directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf := rice.Config{
				LocateOrder: []rice.LocateMethod{
					rice.LocateEmbedded,
					rice.LocateAppended,
					rice.LocateFS,
				},
			}
			box, err := conf.FindBox("boilerplates")
			return nil
		},
	}
	return cmd
}