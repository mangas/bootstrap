package osutil

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"reflect"
	"text/tabwriter"
)

var (
	errColor       = color.New(color.FgRed).SprintFunc()
	infoColor      = color.New(color.FgBlue).SprintFunc()
	stringColor    = color.New(color.FgGreen).SprintFunc()
	delimiterColor = color.New(color.FgBlack).SprintFunc()
	titleColor     = color.New(color.FgYellow).SprintFunc()
)

// Exists returns whether the given file or directory exists or not.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// GetTabWriterOutput gets tabwriter output
func getTabWriterOutput() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 12, 2, 4, ' ', tabwriter.TabIndent)
}

func PrintRows(out *tabwriter.Writer, rows reflect.Value) {
	for i := 0; i < rows.NumField(); i++ {
		if rows.Field(i).String() != "" {
			printRow(
				out, "\n%s\t%s\t%v",
				rows.Type().Field(i).Name,
				rows.Field(i).Interface(),
			)
		}
	}
}

func printRow(out io.Writer, format, key string, value interface{}) {
	if format == "" {
		fmt.Fprintf(out, "\n%s\t%s\t%v", infoColor(key), errColor(":"), stringColor(value))
	}
	fmt.Fprintf(out, format, infoColor(key), errColor(":"), stringColor(value))
}

func PrintDelimiters(out *tabwriter.Writer) {
	delimiter := "====================================================="
	fmt.Fprintf(out, "\n%s\n", delimiterColor(delimiter))
}
