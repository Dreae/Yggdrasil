package daemon

import (
  "log"
  "github.com/hoisie/web"
  "github.com/fsouza/go-dockerclient"
)

func handleGetServerLogs(client *docker.Client)func(*web.Context, string) {
  handler := func(ctx *web.Context, id string) {
    var options docker.LogsOptions
    options.Container = id
    options.OutputStream = ctx
    options.Stdout = true
    options.RawTerminal = true
    options.Tail = "200"
    if err := client.Logs(options); err != nil {
      log.Panicln(err)
    }
  }

  return handler
}
