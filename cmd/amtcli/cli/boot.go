package cli
  
import (
  "log"
  "sync"

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
    
    var wg sync.WaitGroup
    wg.Add(len(args) - 1)

    for _, host := range args[1:] {
      go setBootTarget(host, target, &wg)
    }

    wg.Wait()
  },
}

func setBootTarget(host string, target amt.BootTarget, wg *sync.WaitGroup) {
  defer wg.Done()
  
  client := client(host)
  client.SetBootTarget(target)
  
  log.Println("[" + host + "] Success")
}

func init() {
  rootCmd.AddGroup(bootGroup)
  rootCmd.AddCommand(bootCmd)
}
