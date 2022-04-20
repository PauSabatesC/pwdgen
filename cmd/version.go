package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	VERSION   = "v1.0"
	CODE_INFO = "\ngithub.com/PauSabatesC/\n"
)

var info bool

func init() {
	rootCmd.AddCommand(version)
	version.Flags().BoolVarP(
		&info,
		"all",
		"a",
		false,
		"Get more info about the software",
	)
}

var version = &cobra.Command{
	Use:   "version",
	Short: "Software version",
	Run: func(cmd *cobra.Command, args []string) {
		if info {
			fmt.Println(VERSION, CODE_INFO)
		} else {
			fmt.Println(VERSION)
		}
	},
}
