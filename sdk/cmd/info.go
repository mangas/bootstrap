package cmd

import (
	"github.com/getcouragenow/bootstrap/sdk/pkg/common/osutil"
	"github.com/spf13/cobra"
)

func NewOsInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "prints os info",
		RunE: func(cmd *cobra.Command, args []string) error {
			newUserInfo, err := osutil.InitUserOsEnv()
			if err != nil {
				return err
			}
			if err = newUserInfo.PrintUserOsEnv(); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

