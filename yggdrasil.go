package main

import (
  "fmt"
  "log"
  "flag"
  "io/ioutil"
  "encoding/json"
  "github.com/dreae/yggdrasil/daemon"
  "github.com/dreae/yggdrasil/server"
)

func main() {
  config := readConfig()
  if !config.Server {
    daemon.ServeHttp(config.IP, config.Port, config.JoinURL)
  } else {
    host := fmt.Sprintf("%s:%d", config.HostIP, config.HostPort)
    peers := fmt.Sprintf("%s:%d", config.PeerIP, config.PeerPort)

    server.Listen(host, peers)
  }
}

type Config struct {
  Server bool
  PeerPort int
  PeerIP string
  HostPort int
  HostIP string
  JoinURL string
  IP string
  Port int
}

func readConfig() *Config {
  server_ := flag.Bool("server", false, "Should this instance server as master?")
  peerPort := flag.Int("peer-port", 7315, "Port the master server will listen for daemons on")
  peerIp := flag.String("peer-ip", "127.0.0.1", "Address to listen for daemons on")
  hostPort := flag.Int("host-port", 8080, "Port the master server will server the UI from")
  hostIp := flag.String("host-ip", "0.0.0.0", "Address to serve the UI from")
  configFile := flag.String("config-file", "", "Config file to use")

  port := flag.Int("p", 4315, "Define the listen port")
  ip := flag.String("ip", "0.0.0.0", "IP address to bind to")
  join := flag.String("join-url", "http://127.0.0.1:7315", "Address of the master server")

  flag.Parse()

  if *configFile != "" {
    bytes, err := ioutil.ReadFile(*configFile)
    if err != nil {
      log.Panicln("Unable to read config file: ", err)
    }
    var config map[string]interface{}

    if err = json.Unmarshal(bytes, &config); err != nil {
      log.Panicln("Unable to parse config file: ", err)
    }

    if val, ok := config["server"]; ok {
      *server_ = val.(bool)
    }
    if val, ok := config["peer_port"]; ok {
      *peerPort = int(val.(float64))
    }
    if val, ok := config["peer_ip"]; ok {
      *peerIp = val.(string)
    }
    if val, ok := config["host_port"]; ok {
      *hostPort = int(val.(float64))
    }
    if val, ok := config["host_ip"]; ok {
      *hostIp = val.(string)
    }
    if val, ok := config["join_url"]; ok {
      *join = val.(string)
    }
  }

  var config Config
  config.Server = *server_
  config.PeerPort = *peerPort
  config.PeerIP = *peerIp
  config.HostPort = *hostPort
  config.HostIP = *hostIp
  config.JoinURL = *join
  config.Port = *port
  config.IP = *ip

  return &config
}
