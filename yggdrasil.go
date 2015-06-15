package main

import (
  "fmt"
  "flag"
  "github.com/dreae/yggdrasil/daemon"
)

func main() {
  port := flag.Int("p", 4315, "Define the listen port")
  ip := flag.String("ip", "0.0.0.0", "IP address to bind to")
  flag.Parse()

  daemon.ServeHttp(fmt.Sprintf("%s:%d", *ip, *port))
}
