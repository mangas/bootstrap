package cmd

import (
	"fmt"
	"github.com/getcouragenow/bootstrap/sdk/pkg/statics"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

var (
	namespace string
)

func NewInitBoilerplateCmd() *cobra.Command {
	joinedNS := strings.Join(statikNamespaces, "|")
	usage := fmt.Sprintf("init -n [%s] <output_dir>", joinedNS)
	cmd := &cobra.Command{
		Use:   usage,
		Args:  cobra.ExactArgs(1),
		Short: "Write boilerplates to your current directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println(namespace)
			bp, err := statics.NewBPAsset(statikNamespaces, namespace)
			if err != nil {
				return err
			}
			if err = bp.WriteAllFiles(args[0]); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.PersistentFlags().StringVarP(&namespace, "namespace", "n",
		"bproot", fmt.Sprintf("select boilerplate to unpack, available options are: %s", joinedNS))
	return cmd
}
