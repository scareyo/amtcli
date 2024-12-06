package cli

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/scareyo/amtcli/pkg/amt"
)

var (
  infoCmd = &cobra.Command {
    Use:      "info host ...",
    Short:    "Get device info",
    Run: func(cmd *cobra.Command, args []string) {
      for _, host := range args {
        client := amt.Create(amt.ClientParameters{
          Host: host,
          UseTls: true,
          Username: username,
          Password: password,
        })
        info := client.GetInfo()
        fmt.Println(info.AmtFqdn)
        fmt.Println(info.HostFqdn)
      }
    },
  }
)

func init() {
  rootCmd.AddCommand(infoCmd)
}
