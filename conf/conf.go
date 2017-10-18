package conf

import (
	"time"
)

var (
	// client conf
	WSAddr                  = "ws://127.0.0.1:3653"
	ConnNum                 = 1
	PendingWriteNum         = 2000
	MaxMsgLen        uint32 = 4096
	ConnectInterval         = 3 * time.Second
	HandshakeTimeout        = 10 * time.Second

	// wechat login
	Unioinids []string = []string{
		"o8c-nt6tO8aIBNPoxvXOQTVJUxY0",
		"o8c-ntxW4cW601v6RjaAsExy98w4",
		"o8c-nt2jC5loIHg1BQGgYW6aqe60",
		"o8c-nt6xAZdXAwrQKQ-eIVLr8XRI",
	}
)
