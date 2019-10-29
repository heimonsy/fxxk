package fxxk

type connectClient struct {
	id        string
	done      chan struct{}
	stream    Fxxk_ConnectServer
	idleTunel []tunelStream
}

type tunelStream struct {
	Fxxk_ProxyTunelServer
	done chan struct{}
}
