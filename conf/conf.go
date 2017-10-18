package conf

import (
	"log"
	"time"
)

var (
	// log conf
	LogFlag = log.LstdFlags | log.Lshortfile

	// client conf
	WSAddr                  = "ws://127.0.0.1:3653"
	ConnNum                 = 1
	PendingWriteNum         = 2000
	MaxMsgLen        uint32 = 4096
	ConnectInterval         = 3 * time.Second
	HandshakeTimeout        = 10 * time.Second
)
