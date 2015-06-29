package server

import (
  "log"
  "reflect"
  "database/sql"
  "github.com/hoisie/web"
)

func handleAuthCheck(f interface{}, reqLevel string, conn *sql.DB)interface{} {
  handler := func(ctx *web.Context, args ...interface{}) {
    apiKey := ctx.Request.Header.Get("X-API-Key")
    if apiKey == "" {
      ctx.WriteHeader(401)
      return
    }

    row := conn.QueryRow(
      "SELECT users.username, users.level FROM users, api_keys " +
      "WHERE api_keys.key = $1 AND users.username = api_keys.username",
      apiKey)

    var username string
    var level string
    err := row.Scan(&username, &level)
    switch {
    case err == sql.ErrNoRows:
      ctx.WriteHeader(403)
      return
    case err != nil:
      log.Panicln(err)
    default:
      if level != reqLevel {
        ctx.WriteHeader(403)
        return
      }
    }

    var argv []reflect.Value
    argv = append(argv, reflect.ValueOf(ctx))
    for _, arg := range args {
      argv = append(argv, reflect.ValueOf(arg))
    }

    reflect.ValueOf(f).Call(argv)
  }

  return handler;
}
