package msg


//C2S

type C2S_Heartbeat struct{}

type C2S_WeChatLogin struct {
	Nickname   string
	Headimgurl string
	Sex        int    // 1为男性，2为女性
	Serial     string // 安卓设备硬件序列号,例如:a1113028
	Model      string // 安卓手机型号,例如:MI NOTE Pro
	Unionid    string // 微信unionid
}

//S2C

