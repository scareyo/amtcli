package cli

import (
  "encoding/json"
  "fmt"
  "log"
  "sync"

  "github.com/spf13/cobra"
)

var (
  infoCmd = &cobra.Command {
    Use:      "info host...",
    Short:    "Get device info",
    Run: func(cmd *cobra.Command, args []string) {
      var wg sync.WaitGroup
      wg.Add(len(args))
      for _, host := range args {
        go getInfo(host, &wg)
      }
      wg.Wait()
    },
  }
)

func getInfo(host string, wg *sync.WaitGroup) {
  defer wg.Done()
  client := client(host)
  info, err := json.MarshalIndent(client.GetInfo(), "", "  ")
  if err != nil {
    log.Println(err)
  }

  fmt.Println(string(info))
}

func init() {
  rootCmd.AddCommand(infoCmd)
}
