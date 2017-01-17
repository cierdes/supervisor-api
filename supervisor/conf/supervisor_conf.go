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
	programs []*confProgram
	groups []*confGroup
	include *confInclude
	eventListeners []*confEventListener
	//RpcInterface *svrConfRPCInterface
}

func (c *SupervisorConf)EncodeToString()string{
	var ret string

	if c.unixHttpServer != nil{
		ret += c.unixHttpServer.encode()
		ret += "\n\n"
	}

	if c.inetHttpServer != nil{
		ret += c.inetHttpServer.encode()
		ret += "\n\n"
	}

	if c.supervisord != nil{
		ret += c.supervisord.encode()
		ret += "\n\n"
	}

	if c.supervisorctl != nil{
		ret += c.supervisorctl.encode()
		ret += "\n\n"
	}

	for _, v := range c.programs{
		ret += v.encode()
		ret += "\n\n"
	}

	if c.include != nil{
		ret += c.include.encode()
		ret += "\n\n"
	}

	for _,v := range c.groups{
		ret += v.encode()
		ret += "\n\n"
	}

	for _,v := range c.eventListeners{
		ret += v.encode()
		ret += "\n\n"
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

func (c *SupervisorConf)WriteProgram(name string, key string, val string)error{
	for _,v :=range c.programs{
		if v.name == name{
			return setValue(v.content, key, val)
		}
	}

	p := newConfProgram(name)
	err := setValue(p.content, key, val)
	if err != nil{
		return err
	}

	c.programs = append(c.programs, p)
	return nil
}

func (c *SupervisorConf)WriteInclude(key string, val string)error{
	if c.include == nil{
		c.include = newConfInclude()
	}

	return setValue(c.include.content, key, val)
}

func (c *SupervisorConf)WriteGroup(name string, key string, val string)error{
	for _,v := range c.groups{
		if v.name == name{
			return setValue(v.content, key, val)
		}
	}

	g := newConfGroup(name)
	err := setValue(g.content, key, val)
	if err != nil{
		return err
	}

	c.groups = append(c.groups, g)
	return nil
}

func (c *SupervisorConf)WriteEventListener(name string, key string, val string)error{
	for _,v := range c.eventListeners{
		if v.name == name{
			return setValue(v.content, key, val)
		}
	}

	e := newConfEventListener(name)
	err := setValue(e.content, key, val)
	if err != nil{
		return err
	}

	c.eventListeners = append(c.eventListeners, e)
	return nil
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