package server

import (
  "database/sql"
  "github.com/hoisie/web"
)

func configureUIServer(server *web.Server, conn *sql.DB) {
  server.Config.StaticDir = "ui"
  server.Post("/api/login", handleUserLogin(conn))
  server.Get("/api/servers", handleAuthCheck(handleGetServerList(conn), "admin", conn))
}
