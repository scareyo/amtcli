package cli
  
import (
  "fmt"
  "github.com/spf13/cobra"
)

var powerGroup = &cobra.Group {
  ID:     "power",
  Title:  "Power Commands",
}

var resetCmd = &cobra.Command {
  Use:      "reset host ...",
  Short:    "Reset the device",
  GroupID:  "power",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("[**Unimplemented**] Resetting...")
  },
}

var restartCmd = &cobra.Command {
  Use:      "restart host ...",
  Short:    "Restart the device",
  GroupID:  "power",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("[**Unimplemented**] Restarting...")
  },
}

var onCmd = &cobra.Command {
  Use:      "on host ...",
  Short:    "Power on the device",
  GroupID:  "power",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("[**Unimplemented**] Powering on...")
  },
}

var offCmd = &cobra.Command {
  Use:      "off host ...",
  Short:    "Power off the device",
  GroupID:  "power",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("[**Unimplemented**] Powering off...")
  },
}

func init() {
  rootCmd.AddGroup(powerGroup)
  
  rootCmd.AddCommand(resetCmd)
  rootCmd.AddCommand(restartCmd)
  rootCmd.AddCommand(onCmd)
  rootCmd.AddCommand(offCmd)
}
