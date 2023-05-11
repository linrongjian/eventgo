package gateway

import (
	"cavy/core/app"
	"cavy/core/network/protocols/httpwrap"
	"cavy/core/network/protocols/mqwrap"
)

type GateServer interface {
	app.Server
	Init(...Option) error
	Options() Opts
}

type Option func(*Opts)

type gateServer struct {
	*app.App
	opts Opts
}

func (s *gateServer) Run() error {
	// Ctx = s.Options().Context
	httpwrap.CreateHTTPServer()
	// ClearOnline()
	mqwrap.Startup()
	return nil
}

func (s *gateServer) Stop() error {
	return nil
}

func (s *gateServer) Options() Opts {
	return s.opts
}

func (s *gateServer) Init(opts ...Option) error {
	for _, o := range opts {
		o(&s.opts)
	}
	httpwrap.RegisterGetHandleNoUserID("/", onConnectHandle) //获取入口信息
	s.AddFlags(mqwrap.Flags)
	return nil
}

func (s *gateServer) getOnlineCount() int {
	return 1
}

func (s *gateServer) kickUser(userId string) {
	// return this.wsGateway.kick(userId);
}

func (s *gateServer) broadcast(data interface{}) {
	// return this.wsGateway.broadcast(data);
}

func (s *gateServer) notify(userId string, data interface{}) {
	// return this.wsGateway.notify(userId, data)
}

func NewGateServer(opts ...Option) GateServer {
	options := Opts{}
	for _, o := range opts {
		o(&options)
	}
	app := app.NewApp()
	app.Init()
	return &gateServer{
		App:  app,
		opts: options,
	}
}