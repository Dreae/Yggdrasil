package daemon

import (
  "log"
  "github.com/hoisie/web"
  "github.com/fsouza/go-dockerclient"
)

func handleGetServerFile(client *docker.Client)func(*web.Context, string, string) {
  handler := func(ctx *web.Context, id string, file string) {
    var copyOpts docker.CopyFromContainerOptions
    copyOpts.Container = id
    copyOpts.Resource = file
    copyOpts.OutputStream = ctx
    ctx.SetHeader("Content-Type", "application/x-tar", true)
    if err := client.CopyFromContainer(copyOpts); err != nil {
      log.Panicln(err)
    }
  }

  return handler
}
