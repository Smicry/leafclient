package node

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"leafclient/conf"
	"reflect"
)

var (
	clients []*network.WSClient
)

type Agent struct {
	conn *network.WSConn
}

func Online() {
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
}

func newAgent(conn *network.WSConn) network.Agent {
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

		if err == nil {
			// 处理body
		} else {
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
