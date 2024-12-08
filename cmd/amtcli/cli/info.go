package cli

import (
  "encoding/json"
  "fmt"
  "log"

  "github.com/spf13/cobra"
)

var (
  infoCmd = &cobra.Command {
    Use:      "info host ...",
    Short:    "Get device info",
    Run: func(cmd *cobra.Command, args []string) {
      for _, host := range args {
        client := client(host)
        info, err := json.MarshalIndent(client.GetInfo(), "", "  ")
        if err != nil {
          log.Fatal(err)
        }

        fmt.Println(string(info))
      }
    },
  }
)

func init() {
  rootCmd.AddCommand(infoCmd)
}
