package main

import (
  "os"
  "log"
  "fmt"
  "flag"
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

  port := flag.Int("p", 4315, "Define the listen port")
  ip := flag.String("ip", "0.0.0.0", "IP address to bind to")
  flag.Parse()

  daemon.ServeHttp(fmt.Sprintf("%s:%d", *ip, *port))
}
