package node

import (
	"github.com/name5566/leaf/network"
	"leafclient/conf"
	"github.com/name5566/leaf/log"
	"encoding/json"
	"reflect"
)

var (
	clients []*network.WSClient
)

type Agent struct {
	conn *network.WSConn
}

func Online() {
	for _, uid := range conf.Unioinids {
		client := new(network.WSClient)
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
		log.Debug("client unioinid : %v online", uid)
	}
}

func newAgent(conn *network.WSConn) network.Agent {
	a := new(Agent)
	a.conn = conn
	return a
}

func (a *Agent) Run() {
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
	}
}

func (a *Agent) WriteMsg(msg interface{}) {

	data, err := json.Marshal(msg)
	if err != nil {
		log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
		return
	}
	err = a.conn.WriteMsg(data)
	if err != nil {
		log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
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
