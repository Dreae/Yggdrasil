package server

import (
  "database/sql"
  "github.com/hoisie/web"
)

func configureMasterServer(server *web.Server, conn *sql.DB) {
  server.Post("/join", handleServerJoin(conn))
}
