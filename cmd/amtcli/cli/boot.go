package cli
  
import (
  "log"

  "github.com/scareyo/amtcli/pkg/amt"

  "github.com/spf13/cobra"
)

var bootGroup = &cobra.Group {
  ID:     "boot",
  Title:  "Boot Commands",
}

var bootCmd = &cobra.Command {
  Use:      "boot target host...",
  Short:    "Set the boot target (hdd, cd, pxe, bios)",
  GroupID:  "boot",
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("Changing boot device...")

    var target amt.BootTarget
    switch args[0] {
      case "hdd":
        target = amt.BootTargetHdd
      case "cd":
        target = amt.BootTargetCd
      case "pxe":
        target = amt.BootTargetPxe
      case "bios":
        target = amt.BootTargetBios
    }

    for _, host := range args[1:] {
      client := client(host)
      client.SetBootTarget(target)
    }
  },
}

func init() {
  rootCmd.AddGroup(bootGroup)
  rootCmd.AddCommand(bootCmd)
}
