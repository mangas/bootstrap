package main

import (
	c "github.com/getcouragenow/bootstrap/sdk/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd     *cobra.Command
	globalUsage = `
	bs is a bootstrapper utility for kickstarting development with getcouragenow
	`
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "bs",
		Short:        globalUsage,
		Long:         globalUsage,
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			log.SetLevel(log.Level(1))
		},
	}
	cmd.AddCommand(
		c.NewOsInfoCmd(),
		c.NewInstallToolsCmd(),
		c.NewInitBoilerplateCmd(),
		// newSelfUpgradeCmd(),
	)
	return cmd
}

func main() {
	rootCmd = newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
