package rpchelper

import (
  "log"
  "net"
  "net/rpc"
  "net/rpc/jsonrpc"
)

// ServeRpc starts an RPC server with given host address and service
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

// NewClient creates an RPC client with given host address
func NewClient(host string) (*rpc.Client, error) {
  conn, err := net.Dial("tcp", host)
  if err != nil {
    return nil, err
  }
  return jsonrpc.NewClient(conn), nil
}

// CreateClientPool creates and sends RPC clients to client channel
// It performs load balancing among hosts
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