package cmd

import (
	"fmt"
	"github.com/getcouragenow/bootstrap/sdk/pkg/common/colorutil"
	"github.com/getcouragenow/bootstrap/sdk/pkg/common/osutil"
	"github.com/getcouragenow/bootstrap/sdk/pkg/common/termutil"
	"github.com/spf13/cobra"
	"strings"
)

func NewListToolsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tools available",
		Long:  "List all tools available",
		RunE: func(cmd *cobra.Command, args []string) error {
			contents := termutil.Contents{}
			for k, _ := range toolPathMap {
				installed := osutil.BinExists(strings.TrimSpace(k))
				contents[k] = []string{colorutil.ColorRed(osutil.CrossMark) + " not installed"}
				if installed {
					contents[k] = []string{colorutil.ColorGreen(osutil.CheckMark) + " installed"}
				}
			}
			if _, err := fmt.Println(contents.String("All Tools")); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
