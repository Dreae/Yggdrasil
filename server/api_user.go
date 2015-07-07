package server

import (
  "fmt"
  "io/ioutil"
  "database/sql"
  "encoding/json"
  "github.com/hoisie/web"
  "golang.org/x/crypto/bcrypt"
)

func handleGetServerList(conn *sql.DB)func(*web.Context) {
  handler := func(ctx *web.Context) {
    body, err := ioutil.ReadAll(ctx.Request.Body)
    if err != nil {
      log.Panicln(err)
    }

    var content map[string]string
    err = json.Unmarshal(body, &content)
    if er != nil {
      log.Panicln(err)
    }

    var pwHash string
    row := conn.QueryRow("SELECT password FROM users WHERE username = $1", body["username"])
    err = row.Scan(&pwHash)
    switch {
    case err == sql.ErrNoRows:
      log.Println(fmt.Sprintf("Unknown user %s", content["username"]))
      ctx.WriteHeader(404)
      return
    case err != nil:
      log.Panicln(err)
    default:
      err = bcrypt.CompareHashAndPassword(pwHash, content["password"])
      if err != nil {
        log.Println(fmt.Sprintf("Bad password for user %s", content["username"]))
        ctx.WriteHeader(404)
        return
      }
    }
  }

  return handler;
}
