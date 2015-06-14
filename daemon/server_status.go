package daemon

import (
  "log"
  "fmt"
  "io/ioutil"
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

func handleSetServerStatus(client *docker.Client)func(*web.Context, string) {
  handler := func(ctx *web.Context, id string) {
    body, err := ioutil.ReadAll(ctx.Request.Body)
    if err != nil {
      log.Panicln(err)
    }

    var content map[string]string
    err = json.Unmarshal(body, &content)
    if err != nil {
      ctx.WriteHeader(400)
      return
    }

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

    switch content["State"] {
    case "running":
      if status.State != "running" {
        if status.State == "paused" {
          if err = client.UnpauseContainer(id); err != nil {
            log.Panicln(err)
          }
        } else {
          if err = client.StartContainer(id, nil); err != nil {
            log.Panicln(err)
          }
        }
      }
    case "restarting":
      if status.State == "running" {
        if err = client.RestartContainer(id, 5); err != nil {
          log.Panicln(err)
        }
      }
    case "stopped":
      if status.State != "stopped" {
        if err = client.StopContainer(id, 5); err != nil {
          log.Panicln(err)
        }
      }
    case "paused":
      if status.State == "running" {
        if err = client.PauseContainer(id); err != nil {
          log.Panicln(err)
        }
      }
    }

    ctx.WriteHeader(204)
  }

  return handler
}
