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

    server.Listen(host, peers, config.SQL_DSN)
  }
}

type Config struct {
  Server bool
  PeerPort int
  PeerIP string
  HostPort int
  HostIP string
  SQL_DSN string
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
  sqlDsn := flag.String("sql-dsn", "postgres://root:password@/yggdrasil?sslmode=disable", "DSN to connect to the SQL database")
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
    if val, ok := config["sql_dsn"]; ok {
      *sqlDsn = val.(string)
    }
  }

  return &Config {
    Server: *server_,
    PeerPort: *peerPort,
    PeerIP: *peerIp,
    HostPort: *hostPort,
    HostIP: *hostIp,
    JoinURL: *join,
    Port: *port,
    IP: *ip,
    SQL_DSN: *sqlDsn,
  }
}
