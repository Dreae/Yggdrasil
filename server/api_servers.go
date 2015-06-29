package server

import (
  "database/sql"
  "github.com/hoisie/web"
)

func handleGetServerList(conn *sql.DB)func(*web.Context) {
  handler := func(ctx *web.Context) {
    ctx.WriteString("[]")
  }

  return handler;
}
