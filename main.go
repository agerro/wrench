package main

import (
	"github.com/spf13/cobra"
	"local/wrench/cmd"
	"fmt"
)

func main() {
  cmd.Execute()
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Wrench",
	Long:  `All software has versions. This is Wrench's`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("Wrench v0.1 -- HEAD")
	},
  }