package cmd

import (
	"fmt"
	"os"
	"pwdgen/generator"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().IntP(
		"size",
		"s",
		generator.PASSWORD_DEFAULT_SIZE,
		"Size of the password",
	)
}

var rootCmd = &cobra.Command{
	Use:   "pwdgen [-s size]",
	Short: "pwdgen generates a secure password",
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetInt("size")
		password, err := generator.NewPasswordAllValues(size)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Password ->", password)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
