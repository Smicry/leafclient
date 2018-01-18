package conf

import (
	"time"
)

var (
	// client conf
	WSAddr                  = "ws://192.168.1.34:3661"
	ConnNum                 = 1
	PendingWriteNum         = 2000
	MaxMsgLen        uint32 = 4096
	ConnectInterval         = 3 * time.Second
	HandshakeTimeout        = 10 * time.Second
)
