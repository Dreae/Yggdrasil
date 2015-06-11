package main

import (
  "os"
  "log"
  "github.com/dreae/yggdrasil/daemon"
)

func main() {
  _, err := os.Stat(".yggdrasil")
  if err != nil && os.IsNotExist(err) {
    log.Println("No yggdrasil directory found, creating")
    err := os.Mkdir(".yggdrasil", 0700)
    if err != nil {
      log.Fatalln(err)
    }
  }

  steamErr := daemon.Init_SteamCmd()
  if steamErr != nil {
    log.Fatalln("Error getting steamcmd: ", err)
  }
}
