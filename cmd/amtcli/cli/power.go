package cli

import (
  "log"
  "sync"

  "github.com/scareyo/amtcli/pkg/amt"

  "github.com/spf13/cobra"
)

var powerGroup = &cobra.Group {
  ID:     "power",
  Title:  "Power Commands",
}

var resetCmd = &cobra.Command {
  Use:      "reset host...",
  Short:    "Reset the device",
  GroupID:  "power",
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("Resetting...")

    var wg sync.WaitGroup
    wg.Add(len(args))

    for _, host := range args {
      go setPowerState(host, amt.PowerStateReset, &wg)
    }

    wg.Wait()
  },
}

var restartCmd = &cobra.Command {
  Use:      "restart host...",
  Short:    "Restart the device",
  GroupID:  "power",
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("Restarting...")

    var wg sync.WaitGroup
    wg.Add(len(args))

    for _, host := range args {
      go setPowerState(host, amt.PowerStateRestart, &wg)
    }

    wg.Wait()
  },
}

var onCmd = &cobra.Command {
  Use:      "on host...",
  Short:    "Power on the device",
  GroupID:  "power",
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("Powering on...")

    var wg sync.WaitGroup
    wg.Add(len(args))

    for _, host := range args {
      go setPowerState(host, amt.PowerStateOn, &wg)
    }

    wg.Wait()
  },
}

var offCmd = &cobra.Command {
  Use:      "off host...",
  Short:    "Power off the device",
  GroupID:  "power",
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("Powering off...")

    var wg sync.WaitGroup
    wg.Add(len(args))

    for _, host := range args {
      go setPowerState(host, amt.PowerStateOff, &wg)
    }

    wg.Wait()
  },
}

func setPowerState(host string, state amt.PowerState, wg *sync.WaitGroup) {
  defer wg.Done()
  
  client := client(host)
  success := client.SetPowerState(state)
  
  if !success {
    log.Println("[" + host + "] Failure")
  } else {
    log.Println("[" + host + "] Success")
  }
}

func init() {
  rootCmd.AddGroup(powerGroup)
  
  rootCmd.AddCommand(resetCmd)
  rootCmd.AddCommand(restartCmd)
  rootCmd.AddCommand(onCmd)
  rootCmd.AddCommand(offCmd)
}
