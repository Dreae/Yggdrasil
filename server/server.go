package server

import (
  "log"
  "database/sql"
  "github.com/hoisie/web"
  _ "github.com/lib/pq"
)

func Listen(host string, peers string, sqlDsn string) {
  uiServer := web.NewServer()
  peerServer := web.NewServer()

  conn, err := sql.Open("postgres", sqlDsn)
  if err != nil {
    log.Panicln("Error connecting to database: ", err)
  }


  configureUIServer(uiServer, conn)
  configureMasterServer(peerServer, conn)

  go func() {
    uiServer.Run(host)
  }()

  peerServer.Run(peers)
}
