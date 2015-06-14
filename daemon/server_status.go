package daemon

import (
  "log"
  "fmt"
  "encoding/json"
  "github.com/hoisie/web"
  "github.com/fsouza/go-dockerclient"
)

func handleGetServerStatus(client *docker.Client)func(*web.Context, string) {
  handler := func(ctx *web.Context, id string) {
    status, err := getServerStatus(client, id)
    if err != nil {
      switch err.Error() {
      case fmt.Sprintf("No such container: %s", id):
        ctx.WriteHeader(404)
        return
      default:
        log.Panicln(err)
      }
    }
    ctx.SetHeader("Content-Type", "application/json", true)
    data, err := json.Marshal(status)
    if err != nil {
      log.Panicln(err)
    }
    ctx.WriteString(string(data))
  }

  return handler
}
