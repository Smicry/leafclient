package node

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"leafclient/conf"
	"leafclient/net"
)

var (
	clients []*net.WSClient
)

type Agent struct {
	conn *net.WSConn
}

func Online() {
	client := new(net.WSClient)
	client.Addr = conf.WSAddr
	client.ConnNum = conf.ConnNum
	client.ConnectInterval = conf.ConnectInterval
	client.PendingWriteNum = conf.PendingWriteNum
	client.MaxMsgLen = conf.MaxMsgLen
	client.HandshakeTimeout = conf.HandshakeTimeout
	client.AutoReconnect = false
	client.NewAgent = newAgent

	client.Start()
	clients = append(clients, client)
}

func newAgent(conn *net.WSConn) net.Agent {
	a := new(Agent)
	a.conn = conn
	return a
}

func (a *Agent) Run() {
	a.ReadMsg()
}

func (a *Agent) ReadMsg() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}

		var body interface{}
		err = json.Unmarshal(data, body)

		if err != nil {
			log.Debug("unmarshal message error: %v", err)
			break
		}

		log.Debug("readMsg: %+v", body)
	}
}

func (a *Agent) WriteMsg(msg interface{}) {
	err := a.conn.WriteMsg(msg)
	if err != nil {
		log.Debug("write message: %v", err)
	}
}

func (a *Agent) OnClose() {
	a.conn.Close()
}

func Destroy() {
	for _, client := range clients {
		client.Close()
	}
}
