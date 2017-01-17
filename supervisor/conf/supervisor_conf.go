/*
@author: foolbread
@time: 2017-01-17
@file: supervisor-api/supervisor/conf/supervisor_conf.go
*/

package conf

import (
	"errors"
)

type SupervisorConf struct {
	unixHttpServer *confUnixHttpServer
	inetHttpServer *confInetHttpServer
	supervisord    *confSupervisord
	supervisorctl *confSupervisorctl
	//Programs []*svrConfFcgiProgram
	//Groups []*svrConfGroup
	//EventListener *svrConfEventListener
	//RpcInterface *svrConfRPCInterface
}

func (c *SupervisorConf)EncodeToString()string{
	var ret string

	if c.unixHttpServer != nil{
		ret += c.unixHttpServer.encode()
	}

	if c.inetHttpServer != nil{
		ret += c.inetHttpServer.encode()
	}

	if c.supervisord != nil{
		ret += c.supervisord.encode()
	}

	if c.supervisorctl != nil{
		ret += c.supervisorctl.encode()
	}

	return ret
}

func (c *SupervisorConf)WriteUnixHttpServer(key string, val string)error{
	if c.unixHttpServer == nil{
		c.unixHttpServer = newConfUnixHttpServer()
	}

	return setValue(c.unixHttpServer.content, key, val)
}

func (c *SupervisorConf)WriteInetHttpServer(key string, val string)error{
	if c.inetHttpServer == nil{
		c.inetHttpServer = newConfInetHttpServer()
	}

	return setValue(c.inetHttpServer.content, key, val)
}

func (c *SupervisorConf)WriteSupervisord(key string, val string)error{
	if c.supervisord == nil{
		c.supervisord = newConfSupervisord()
	}

	return setValue(c.supervisord.content, key, val)
}

func (c *SupervisorConf)WriteSupervisorctl(key string, val string)error{
	if c.supervisorctl == nil{
		c.supervisorctl = newConfSupervisorctl()
	}

	return setValue(c.supervisorctl.content, key, val)
}

//////////////////////////////////////////////////////////////////////////////////////////
func setValue(data map[string]string, key string, val string)error{
	_,ok := data[key]
	if !ok {
		return errors.New("key is invalid!")
	}

	data[key] = val
	return nil
}