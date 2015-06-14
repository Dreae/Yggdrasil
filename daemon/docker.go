package daemon

import (
  "os"
  "fmt"
  "time"
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

type ServerStatus struct {
  State string
  ExitCode int
  Pid int
  Error string
  StartedAt time.Time
  FinishedAt time.Time
}

func createServer(client *docker.Client, def ServerDefinition) error {
  var createConfig docker.CreateContainerOptions
  createConfig.Name = def.Id

  var contConfig docker.Config
  var hostConfig docker.HostConfig

  contConfig.Image = def.Image
  contConfig.Cmd = def.Args
  contConfig.AttachStdin = true
  contConfig.Tty = true

  hostConfig.PortBindings = make(map[docker.Port][]docker.PortBinding)
  contConfig.ExposedPorts = make(map[docker.Port]struct{})

  for _, port := range def.Ports {
    pK := docker.Port(fmt.Sprintf("%d/%s", port.Container, port.Protocol))
    hostConfig.PortBindings[pK] = make([]docker.PortBinding, 1)
    hostConfig.PortBindings[pK][0].HostIP = "0.0.0.0"
    hostConfig.PortBindings[pK][0].HostPort = strconv.Itoa(port.Host)

    contConfig.ExposedPorts[pK] = struct{}{}
  }

  createConfig.Config = &contConfig
  createConfig.HostConfig = &hostConfig
  _, err := client.CreateContainer(createConfig)
  return err
}

func getServerStatus(client *docker.Client, id string) (*ServerStatus, error) {
  container, err := client.InspectContainer(id)
  if err != nil {
    return nil, err
  }

  var status ServerStatus
  switch {
  case container.State.Running && !container.State.Paused:
    status.State = "running"
  case container.State.Paused:
    status.State = "paused"
  case container.State.Restarting:
    status.State = "restarting"
  case true:
    status.State = "stopped"
  }
  status.ExitCode = container.State.ExitCode
  status.Pid = container.State.Pid
  status.Error = container.State.Error
  status.StartedAt = container.State.StartedAt
  status.FinishedAt = container.State.FinishedAt

  return &status, nil
}

func initDockerClient()(*docker.Client, error) {
  var client *docker.Client
  var err error

  endpoint := os.Getenv("DOCKER_HOST")
  if endpoint == "" {
    endpoint = "unix:///var/run/docker.sock"
  }

  path := os.Getenv("DOCKER_CERT_PATH")
  if path == "" {
    client, err = docker.NewClient(endpoint)
  } else {
    ca := fmt.Sprintf("%s/ca.pem", path)
    cert := fmt.Sprintf("%s/cert.pem", path)
    key := fmt.Sprintf("%s/key.pem", path)
    client, err = docker.NewTLSClient(endpoint, cert, key, ca)
  }
  return client, err
}
