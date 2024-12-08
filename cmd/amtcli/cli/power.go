package cli

import (
  "log"

  "github.com/scareyo/amtcli/pkg/amt"

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
    log.Println("Resetting...")
    for _, host := range args {
      client := client(host)
      client.SetPowerState(amt.PowerStateReset)
    }
  },
}

var restartCmd = &cobra.Command {
  Use:      "restart host ...",
  Short:    "Restart the device",
  GroupID:  "power",
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("Restarting...")
    for _, host := range args {
      client := client(host)
      client.SetPowerState(amt.PowerStateRestart)
    }
  },
}

var onCmd = &cobra.Command {
  Use:      "on host ...",
  Short:    "Power on the device",
  GroupID:  "power",
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("Powering on...")
    for _, host := range args {
      client := client(host)
      client.SetPowerState(amt.PowerStateOn)
    }
  },
}

var offCmd = &cobra.Command {
  Use:      "off host ...",
  Short:    "Power off the device",
  GroupID:  "power",
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("Powering off...")
    for _, host := range args {
      client := client(host)
      client.SetPowerState(amt.PowerStateOff)
    }
  },
}

func init() {
  rootCmd.AddGroup(powerGroup)
  
  rootCmd.AddCommand(resetCmd)
  rootCmd.AddCommand(restartCmd)
  rootCmd.AddCommand(onCmd)
  rootCmd.AddCommand(offCmd)
}
