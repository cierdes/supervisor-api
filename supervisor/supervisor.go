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

func (s *Supervisor)GetAPIVersion()(string,error){
	var ret string
	err := s.client.Call("supervisor.getAPIVersion",nil,&ret)

	return ret,err
}

func (s *Supervisor)GetSupervisorVersion()(string,error){
	var ret string
	err := s.client.Call("supervisor.getSupervisorVersion",nil,&ret)

	return ret,err
}

func (s *Supervisor)GetIdentification()(string,error){
	var ret string
	err := s.client.Call("supervisor.getIdentification",nil,&ret)

	return ret,err
}

func (s *Supervisor)GetState()(*SupervisorStatus,error){
	var ret *SupervisorStatus = new(SupervisorStatus)
	err := s.client.Call("supervisor.getState", nil, ret)

	return ret,err
}

func (s *Supervisor)GetPID()(int,error){
	var ret int
	err := s.client.Call("supervisor.getPID", nil, &ret)

	return ret,err
}

func (s *Supervisor)Restart()(bool,error){
	var ret bool
	err := s.client.Call("supervisor.restart", nil, &ret)

	return ret,err
}

//add change remove
//note:经过测试实际上并没有reload远程的config
func (s *Supervisor)ReloadConfig()([][][]string,error){
	var ret [][][]string
	err := s.client.Call("supervisor.reloadConfig",nil,&ret)

	return ret,err
}

func (s *Supervisor)ShutDown()(bool,error){
	var ret bool
	err := s.client.Call("supervisor.shutdown", nil, &ret)

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

//首先要先停掉该进程组，这个进程组才能删除成功
func (s *Supervisor)RemoveProcessGroup(name string)(bool,error){
	var ret bool
	err := s.client.Call("supervisor.removeProcessGroup", name, &ret)

	return ret,err
}

func (s *Supervisor)AddProcessGroup(name string)(bool,error){
	var ret bool
	err := s.client.Call("supervisor.addProcessGroup", name, &ret)

	return ret,err
}

func (s *Supervisor)StopAllProcesses(wait bool)([]SupervisorProcessStatus,error){
	var ret []SupervisorProcessStatus
	err := s.client.Call("supervisor.stopAllProcesses", wait, &ret)

	return ret,err
}
