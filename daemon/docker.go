package daemon

import (
  "fmt"
  "log"
  "strconv"
  "github.com/fsouza/go-dockerclient"
)

type ServerDefinition struct {
  Id string
  Image string
  Ports []PortDefinition
  Args []string
}

type PortDefinition struct {
  Container int
  Host int
  Protocol string
}

func createServer(client *docker.Client, def ServerDefinition) error {
  var createConfig docker.CreateContainerOptions
  createConfig.Name = def.Id

  var contConfig docker.Config
  var hostConfig docker.HostConfig

  contConfig.Image = def.Image
  contConfig.Cmd = def.Args
  contConfig.Tty = true

  hostConfig.PortBindings = make(map[docker.Port][]docker.PortBinding)

  for _, port := range def.Ports {
    pK := docker.Port(fmt.Sprintf("%d/%s", port.Container, port.Protocol))
    hostConfig.PortBindings[pK] = make([]docker.PortBinding, 1)
    hostConfig.PortBindings[pK][0].HostIP = "0.0.0.0"
    hostConfig.PortBindings[pK][0].HostPort = strconv.Itoa(port.Host)
  }

  log.Println(hostConfig.PortBindings)

  createConfig.Config = &contConfig
  createConfig.HostConfig = &hostConfig
  _, err := client.CreateContainer(createConfig)
  return err
}
