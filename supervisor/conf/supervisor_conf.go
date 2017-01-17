package supervisor

type SupervisorConf struct {
	UnixHttpServer *svrConfUnixHttpServer
	InetHttpServer *svrConfInetHttpServer
	Supervisord *svrConfSupervisord
	Supervisorctl *svrConfSupervisorctl
	Programs []*svrConfFcgiProgram
	Groups []*svrConfGroup
	EventListener *svrConfEventListener
	RpcInterface *svrConfRPCInterface
}

var (
	unxitHttpServer []string = []string{"file","chmod","chown","username","password"}
)

type svrConfUnixHttpServer struct {
	content map[string]string
}

func newSvrConfUnixHttpServer()*svrConfUnixHttpServer{
	r := new(svrConfUnixHttpServer)
	r.content = make(map[string]string)
	for _,v := range unxitHttpServer{
		r.content[v] = ""
	}

	return r
}

type svrConfInetHttpServer struct {

}

type svrConfSupervisord struct {

}

type svrConfSupervisorctl struct {

}

type svrConfInclude struct {

}

type svrConfProgram struct {

}

type svrConfGroup struct {

}

type svrConfEventListener struct {

}

type svrConfRPCInterface struct {

}

type svrConfFcgiProgram struct {

}