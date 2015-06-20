package daemon

import (
  "log"
  "github.com/hoisie/web"
)

func ServeHttp(bind string) {
  client, err := initDockerClient()
  if err != nil {
    log.Panicln(err)
  }

  web.Put("/servers/([A-z0-9]+)", handleCreateServer(client))

  web.Get("/servers/([A-z0-9]+)/status", handleGetServerStatus(client))
  web.Put("/servers/([A-z0-9]+)/status", handleSetServerStatus(client))

  web.Get("/servers/([A-z0-9]+)/logs", handleGetServerLogs(client))

  web.Get("/servers/([A-z0-9]+)/files/(.*)", handleGetServerFile(client))

  web.Run(bind)
}
