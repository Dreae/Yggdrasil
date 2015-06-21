package daemon

import (
  "log"
  "fmt"
  "bytes"
  "net/http"
  "encoding/json"
  "github.com/hoisie/web"
)

func ServeHttp(ip string, port int, join string) {
  client, err := initDockerClient()
  if err != nil {
    log.Panicln(err)
  }

  web.Put("/servers/([A-z0-9]+)", handleCreateServer(client))

  web.Get("/servers/([A-z0-9]+)/status", handleGetServerStatus(client))
  web.Put("/servers/([A-z0-9]+)/status", handleSetServerStatus(client))

  web.Get("/servers/([A-z0-9]+)/logs", handleGetServerLogs(client))

  web.Get("/servers/([A-z0-9]+)/files/(.*)", handleGetServerFile(client))

  info := make(map[string]interface{})
  info["ID"] = ""
  info["IP"] = ip
  info["Port"] = port
  body, _ := json.Marshal(info)

  _, err = http.Post(fmt.Sprintf("%s/join", join), "application/json", bytes.NewReader(body))
  if err != nil {
    log.Print("Error joining master server: ")
    log.Panicln(err)
  }

  web.Run(fmt.Sprintf("%s:%d", ip, port))
}
