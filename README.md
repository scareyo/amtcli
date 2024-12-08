# amtcli

A command-line tool for interacting with Intel AMT devices

```
Usage:
  amtcli [command]

Boot Commands
  boot        Set the boot target (hdd, cd, pxe, bios)

Power Commands
  off         Power off the device
  on          Power on the device
  reset       Reset the device
  restart     Restart the device

Additional Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  info        Get device info
  version     Print the version

Flags:
  -h, --help              help for amtcli
  -p, --password string   Intel AMT password
  -u, --username string   Intel AMT username

Use "amtcli [command] --help" for more information about a command.
```

## Environment Variables

`AMTCLI_PASSWORD`: Used for password if `--password` is unset
