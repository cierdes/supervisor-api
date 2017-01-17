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
	UnixHttpServer *confUnixHttpServer
	InetHttpServer *confInetHttpServer
	//Supervisord *svrConfSupervisord
	//Supervisorctl *svrConfSupervisorctl
	//Programs []*svrConfFcgiProgram
	//Groups []*svrConfGroup
	//EventListener *svrConfEventListener
	//RpcInterface *svrConfRPCInterface
}

func (c *SupervisorConf)EncodeToString()string{
	var ret string

	if c.UnixHttpServer != nil{
		ret += c.UnixHttpServer.encode()
	}

	if c.InetHttpServer != nil{
		ret += c.InetHttpServer.encode()
	}

	return ret
}

func (c *SupervisorConf)WriteUnixHttpServer(key string, val string)error{
	if c.UnixHttpServer == nil{
		c.UnixHttpServer = newConfUnixHttpServer()
	}

	return setValue(c.UnixHttpServer.content, key, val)
}

func (c *SupervisorConf)WriteInetHttpServer(key string, val string)error{
	if c.InetHttpServer == nil{
		c.InetHttpServer = newConfInetHttpServer()
	}

	return setValue(c.InetHttpServer.content, key, val)
}

//////////////////////////////////////////////////////////////////////////////////////////
func encode(data map[string]string)[]string{
	var ret []string
	for k,v := range data{
		if len(v) > 0{
			ret = append(ret, k+"="+v)
		}
	}

	return ret
}

func setValue(data map[string]string, key string, val string)error{
	_,ok := data[key]
	if !ok {
		return errors.New("key is invalid!")
	}

	data[key] = val
	return nil
}