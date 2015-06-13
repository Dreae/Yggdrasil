package daemon

import (
  "os"
  "fmt"
  "github.com/hoisie/web"
  "github.com/fsouza/go-dockerclient"
)

func ServeHttp(bind string) {
  var client *docker.Client
  endpoint := os.Getenv("DOCKER_HOST")
  if endpoint == "" {
    endpoint = "unix:///var/run/docker.sock"
  }

  path := os.Getenv("DOCKER_CERT_PATH")
  if path == "" {
    client, _ = docker.NewClient(endpoint)
  } else {
    ca := fmt.Sprintf("%s/ca.pem", path)
    cert := fmt.Sprintf("%s/cert.pem", path)
    key := fmt.Sprintf("%s/key.pem", path)
    client, _ = docker.NewTLSClient(endpoint, cert, key, ca)
  }

  web.Put("/servers/([A-z0-9]+)", handleCreateServer(client))

  web.Run(bind)
}
