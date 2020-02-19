package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	template    *string
	outFileName *string
	languages   *string
	outDir      *string
	filesDir    *string
	inExt       *string
	outExt      *string
	full        *bool
)

// i18nCmd represents the i18n command
var i18nCmd = &cobra.Command{
	Use:   "i18n",
	Short: "Generate json and arb translated files ",
	Long:  `You can generate json and arb translated files in any languages.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if *template == "" {
			return services.GenerateMultiLanguagesFilesFromFiles(*outDir, *filesDir, *inExt, *outExt, *full)
		}
		return services.GenerateMultiLanguageFilesFromTemplate(*template, *outDir, *outFileName, *outExt, "==", getLanguages(*languages, ","), *full)
	},
}

func init() {
	rootCmd.AddCommand(i18nCmd)
	outDir = flutterCmd.Flags().StringP("outDir", "p", ".", "Out dir.")
	filesDir = flutterCmd.Flags().StringP("filesDir", "d", ".", "files directory.")
	template = flutterCmd.Flags().StringP("template", "t", "", "Template path.")
	outFileName = flutterCmd.Flags().StringP("fileName", "n", "", "The out file name.")
	languages = flutterCmd.Flags().StringP("languages", "l", "en,fr,es,de", "The languages generated.")
	inExt = flutterCmd.Flags().StringP("inExt", "i", "json", "Files extension as input can be: \"json\" or \"arb\".")
	outExt = flutterCmd.Flags().StringP("outExt", "o", "arb", "Files extension as output can be: \"json\" or \"arb\".")
	full = flutterCmd.Flags().BoolP("full", "f", false, "Get full detailed out file.")
}

func getLanguages(languages, sep string) []string {
	return strings.Split(languages, sep)
}
