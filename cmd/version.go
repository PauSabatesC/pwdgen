package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionTag string //Will be set automatically when deployment occurs :)

const (
	CODE_INFO = "\ngithub.com/PauSabatesC/"
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
			fmt.Println(versionTag, CODE_INFO)
		} else {
			fmt.Println(versionTag)
		}
	},
}
