package cli

import (
  "fmt"
  "github.com/spf13/cobra"
)

var (
  username string
  password string
)

var (
  rootCmd = &cobra.Command {
    Use:    "amtcli",
    Short:  "A command-line tool for interacting with Intel AMT devices",
  }

  versionCmd = &cobra.Command {
    Use:    "version",
    Short:  "Print the version",
    DisableFlagsInUseLine: true,

    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("amtgo v0.1.0")
    },
  }
)

func init() {
  rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Intel AMT username")
  rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Intel AMT password")

  rootCmd.AddCommand(versionCmd)
}

func Execute() error {
  return rootCmd.Execute()
}

