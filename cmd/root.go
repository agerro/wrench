package cmd

import "github.com/spf13/cobra"
import "fmt"
import "os"

var rootCmd = &cobra.Command{
	Use:   "wrench",
	Short: "wrench is a over-all usable cli tool there to help you build one yourself",
	Long: `tl;dr`,
	Run: func(cmd *cobra.Command, args []string) {
	  // Do Stuff Here
	},
  }

  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }