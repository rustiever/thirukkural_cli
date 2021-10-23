package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	apiKey string = ""
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tl",
	Short: "Get your thirukkrals at commandline",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Thirukkural")
		fmt.Println("Please run ./tl get [1-1330] to see your favourite kural")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
