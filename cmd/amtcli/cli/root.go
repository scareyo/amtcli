package cli

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
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
  var (
    username string
    password string
  )

  viper.SetEnvPrefix("amtcli")

  // Username
  rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Intel AMT username")
  viper.SetDefault("username", "admin")
  viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))

  // Password
  rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Intel AMT password")
  viper.BindEnv("password")
  viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))

  rootCmd.AddCommand(versionCmd)
}

func Execute() error {
  return rootCmd.Execute()
}

