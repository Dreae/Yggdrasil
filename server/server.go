package server

import (
  "github.com/hoisie/web"
)

func Listen(host string, peers string) {
  uiServer := web.NewServer()
  peerServer := web.NewServer()

  go func() {
    uiServer.Run(host)
  }()

  peerServer.Run(peers)
}
