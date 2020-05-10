package cmd

import (
	"errors"
	"fmt"
	"github.com/getcouragenow/bootstrap/sdk/pkg/common/mkutil"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	short = `install specified tool to User's PATH, see 'list' command for list of tools`
	long  = `
	install usage:
		-i <tool_name> to install tool.
		-u <tool_name> to uninstall tool.
		-p <tool_name> to print tool's makefile information.
		-l <tool_name> to list tool's makefile targets.
		-t <tool_name> to run golang test tool, by default it will also run when you install tool.
	`
)

var (
	installOpt, testOpt, helpOpt, printOpt, uninstallOpt bool
)

func NewInstallToolsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tool -[i|u|h|p|t|l] <tool_name>",
		Short: short,
		Long:  long,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			toolPath := toolPathMap[args[0]]
			log.Println(toolPath)
			if toolPath == "" {
				return errors.New(fmt.Sprintf("we're sorry, %s tool does not exists", args[0]))
			}
			if installOpt {
				if err := mkutil.MakeTestTool(toolPath); err != nil {
					return err
				}
				if err := mkutil.MakeInstallTool(toolPath); err != nil {
					return err
				}
			}
			if testOpt {
				if err := mkutil.MakeTestTool(toolPath); err != nil {
					return err
				}
			}
			if helpOpt {
				if err := mkutil.MakeHelpTool(toolPath); err != nil {
					return err
				}
			}
			if printOpt {
				if err := mkutil.MakePrintTool(toolPath); err != nil {
					return err
				}
			}
			if uninstallOpt {
				if err := mkutil.MakeUninstallTool(toolPath); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.PersistentFlags().BoolVarP(
		&installOpt, "install", "i", false, `install tool`,
	)
	cmd.PersistentFlags().BoolVarP(
		&testOpt, "test", "t", false, `test tool`,
	)
	cmd.PersistentFlags().BoolVarP(
		&helpOpt, "list", "l", true, `print tool's makefile targets'`,
	)
	cmd.PersistentFlags().BoolVarP(
		&printOpt, "print", "p", false, `print tool's information'`,
	)
	cmd.PersistentFlags().BoolVarP(
		&uninstallOpt, "uninstall", "u", false, `uninstall tool`,
	)
	return cmd
}
