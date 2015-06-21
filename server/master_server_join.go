package server

import (
  "log"
  "io/ioutil"
  "crypto/sha1"
  "crypto/rand"
  "database/sql"
  "encoding/hex"
  "encoding/json"
  "github.com/hoisie/web"
)

type DaemonServer struct {
  ID string
  IP string
  Port int
}

func handleServerJoin(conn *sql.DB)func(*web.Context) {
  handler := func(ctx *web.Context) {
    body, err := ioutil.ReadAll(ctx.Request.Body)
    if err != nil {
      log.Panicln(err)
    }
    var daemon DaemonServer
    if err = json.Unmarshal(body, &daemon); err != nil {
      log.Panicln(err)
    }

    var dbDaemon DaemonServer
    log.Printf("Searching for dameon '%s'", daemon.ID)
    row := conn.QueryRow("SELECT id, ip, port FROM daemon_servers WHERE id = $1", daemon.ID)
    err = row.Scan(&dbDaemon.ID, &dbDaemon.IP, &dbDaemon.Port)
    switch {
    case err == sql.ErrNoRows:
      r := make([]byte, 16)
      if _, rErr := rand.Read(r); rErr != nil {
        log.Panicln(rErr)
      }

      h := sha1.New()
      daemon.ID = hex.EncodeToString(h.Sum(r))
      _, dbErr := conn.Exec("INSERT INTO daemon_servers VALUES ($1, $2, $3)", daemon.ID, daemon.IP, daemon.Port)
      if dbErr != nil {
        log.Panicln(dbErr)
      }
      log.Printf("Assigned ID %s to new Daemon %s:%d\n", daemon.ID, daemon.IP, daemon.Port)
    case err != nil:
      log.Panicln(err)
    default:
      log.Printf("Reconnected to Daemon %s\n", daemon.ID)
    }
    ctx.WriteHeader(200)
    ctx.WriteString(daemon.ID)
  }

  return handler
}
