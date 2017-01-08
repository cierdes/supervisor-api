/*
@author: foolbread
@time: 2017-01-06
@file: supervisor-api/supervisor/supervisor.go
*/
package supervisor

import (
	"github.com/kolo/xmlrpc"
	"fmt"
)

const url_fmt = "http://%s/RPC2"

type Supervisor struct {
	addr string
	rpcUrl string
	client *xmlrpc.Client
}

func NewSupervisor(addr string)(*Supervisor,error){
	r := new(Supervisor)
	r.addr = addr
	r.rpcUrl = fmt.Sprintf(url_fmt, addr)

	var err error
	r.client,err =xmlrpc.NewClient(r.rpcUrl,nil)
	if err != nil{
		return nil,err
	}

	return r,nil
}

func (s *Supervisor)Close(){
	if s.client != nil{
		s.client.Close()
	}
}

func (s *Supervisor)ListMethods()([]string,error){
	var ret []string
	err := s.client.Call("system.listMethods",nil, &ret)

	return ret,err
}

func (s *Supervisor)MethodHelp(method string)(string,error){
	var ret string
	err := s.client.Call("system.methodHelp", method, &ret)

	return ret,err
}

func (s *Supervisor)GetProcessInfo(name string)(*SupervisorProcessInfo,error){
	var ret *SupervisorProcessInfo = new(SupervisorProcessInfo)
	err := s.client.Call("supervisor.getProcessInfo", name, ret)

	return ret,err
}

func (s *Supervisor)GetAllProcessInfo()([]SupervisorProcessInfo,error){
	var ret []SupervisorProcessInfo
	err := s.client.Call("supervisor.getAllProcessInfo", nil, &ret)

	return ret,err
}

func (s *Supervisor)StartProcess(name string, wait bool)(bool,error){
	var ret bool
	var args []interface{} = []interface{}{name, wait}
	err := s.client.Call("supervisor.startProcess", args, &ret)

	return ret,err
}

func (s *Supervisor)StartAllProcess(wait bool)([]SupervisorProcessStatus,error){
	var ret []SupervisorProcessStatus
	err := s.client.Call("supervisor.startAllProcesses", wait, &ret)

	return ret,err
}

func (s *Supervisor)StartProcessGroup(name string, wait bool)([]SupervisorProcessStatus,error){
	var ret []SupervisorProcessStatus
	var args []interface{} = []interface{}{name, wait}
	err := s.client.Call("supervisor.startProcessGroup", args, &ret)

	return ret,err
}

func (s *Supervisor)StopProcess(name string, wait bool)(bool,error){
	var ret bool
	var args []interface{} = []interface{}{name, wait}
	err := s.client.Call("supervisor.stopProcess", args, &ret)

	return ret,err
}

func (s *Supervisor)StopProcessGourp(name string, wait bool)([]SupervisorProcessStatus,error){
	var ret []SupervisorProcessStatus
	var args []interface{} = []interface{}{name, wait}
	err := s.client.Call("supervisor.stopProcessGroup", args, &ret)

	return ret,err
}

func (s *Supervisor)StopAllProcesses(wait bool)([]SupervisorProcessStatus,error){
	var ret []SupervisorProcessStatus
	err := s.client.Call("supervisor.stopAllProcesses", wait, &ret)

	return ret,err
}

func (s *Supervisor)ReloadConfig()(bool,error){
	var ret bool
	err := s.client.Call("supervisor.reloadConfig", nil, &ret)

	return ret,err
}