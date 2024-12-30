package main

import (
  "log"
  "github.com/scareyo/amtcli/cmd/amtcli/cli"
)

func main() {
  err := cli.Execute()
  if err != nil {
    log.Println(err)
  }
}
