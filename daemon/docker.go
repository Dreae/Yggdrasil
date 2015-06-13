package daemon

import (
  "fmt"
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
  contConfig.Entrypoint = def.Args
  contConfig.Tty = true

  hostConfig.PortBindings = make(map[docker.Port][]docker.PortBinding)

  for _, port := range def.Ports {
    pK := docker.Port(fmt.Sprintf("%d/%s", port.Container, port.Protocol))
    hostConfig.PortBindings[pK] = make([]docker.PortBinding, 1)
    hostConfig.PortBindings[pK][0].HostPort = string(port.Host)
  }


  createConfig.Config = &contConfig
  createConfig.HostConfig = &hostConfig
  _, err := client.CreateContainer(createConfig)
  return err
}
