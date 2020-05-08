package rpchelper

import (
  "log"
  "net"
  "net/rpc"
  "net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
  err := rpc.Register(service)
  if err != nil {
    return err
  }

  listener, err := net.Listen("tcp", host)
  if err != nil {
    return err
  }
  log.Printf("Listening on %s", host)

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Printf("accept error: %v", err)
      continue
    }
    go jsonrpc.ServeConn(conn)
  }
}

func NewClient(host string) (*rpc.Client, error) {
  conn, err := net.Dial("tcp", host)
  if err != nil {
    return nil, err
  }
  return jsonrpc.NewClient(conn), nil
}

func CreateClientPool(hosts []string) chan *rpc.Client {
  var clients []*rpc.Client
  for _, h := range hosts {
    client, err := NewClient(h)
    if err == nil {
      clients = append(clients, client)
    } else {
      log.Printf("error connecting to %s: %v", h, err)
    }
  }

  out := make(chan *rpc.Client)
  go func() {
    for {
      for _, client := range clients {
        out <- client
      }
    }
  }()
  return out
}