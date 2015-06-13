package daemon

import (
  "log"
  "io/ioutil"
  "encoding/json"
  "github.com/hoisie/web"
  "github.com/fsouza/go-dockerclient"
)

func handleCreateServer(client *docker.Client)interface {} {
  handler := func(ctx *web.Context, id string) {
    body, err := ioutil.ReadAll(ctx.Request.Body)
    if err != nil {
      log.Panicln(err)
    }

    var content ServerDefinition
    err = json.Unmarshal(body, &content)
    if err != nil {
      log.Panicln(err)
    }
    content.Id = id
    err = createServer(client, content)
    if err != nil {
      log.Panicln(err)
    }

    ctx.WriteHeader(201)
  }

  return handler
}
