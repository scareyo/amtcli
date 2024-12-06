package cli
  
import (
  "fmt"
  "github.com/spf13/cobra"
)

var bootGroup = &cobra.Group {
  ID:     "boot",
  Title:  "Boot Commands",
}

var bootCmd = &cobra.Command {
  Use:      "boot",
  Short:    "Restart the device",
  GroupID:  "boot",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("[**Unimplemented**] Changing boot device...")
  },
}

func init() {
  rootCmd.AddGroup(bootGroup)
}
