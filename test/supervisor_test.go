/*
@author: foolbread
@time: 2017-01-06
@file: supervisor-api/test/supervisor_test.go
*/
package test

import (
	"testing"

	"github.com/cierdes/supervisor-api/supervisor"
)

func Test_ListMethod(t *testing.T) {
	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r, err := client.ListMethods()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(r)
}

func Test_MethodHelp(t *testing.T) {
	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r, err := client.MethodHelp("system.multicall")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(r)
}

func Test_GetProcessInfo(t *testing.T) {
	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r, err := client.GetProcessInfo("cesi")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("processname:", r.ProcessName)
	t.Log("group:", r.Group)
	t.Log("description:", r.Description)
	t.Log("start:", r.Start)
	t.Log("stop:", r.Stop)
	t.Log("now:", r.Now)
	t.Log("state:", r.State)
	t.Log("statename:", r.StateName)
	t.Log("spawnerr:", r.Spawnerr)
	t.Log("exitstatus:", r.ExitStatus)
	t.Log("logfile:", r.LogFile)
	t.Log("stdout_logfile:", r.StdoutLogFile)
	t.Log("stderr_logfile:", r.StderrLogFile)
	t.Log("pid:", r.Pid)
}

func Test_GetAllProcessInfo(t *testing.T) {
	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r, err := client.GetAllProcessInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(r))

	for _, v := range r {
		t.Log("processname:", v.ProcessName)
		t.Log("group:", v.Group)
		t.Log("description:", v.Description)
		t.Log("start:", v.Start)
		t.Log("stop:", v.Stop)
		t.Log("now:", v.Now)
		t.Log("state:", v.State)
		t.Log("statename:", v.StateName)
		t.Log("spawnerr:", v.Spawnerr)
		t.Log("exitstatus:", v.ExitStatus)
		t.Log("logfile:", v.LogFile)
		t.Log("stdout_logfile:", v.StdoutLogFile)
		t.Log("stderr_logfile:", v.StderrLogFile)
		t.Log("pid:", v.Pid)
	}
}

func Test_StartProcess(t *testing.T) {
	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r, err := client.StartProcess("cesi", true)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("start status:", r)
}

func Test_StartAllProcesses(t *testing.T) {
	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r, err := client.StartAllProcess(true)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range r {
		t.Log("processname:", v.ProcessName)
		t.Log("group:", v.Group)
		t.Log("description:", v.Description)
		t.Log("status:", v.Status)
	}
}

func Test_StartProcessGroup(t *testing.T) {
	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r, err := client.StartProcessGroup("cesi", true)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range r {
		t.Log("processname:", v.ProcessName)
		t.Log("group:", v.Group)
		t.Log("description:", v.Description)
		t.Log("status:", v.Status)
	}
}

func Test_StopProcess(t *testing.T) {
	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r, err := client.StopProcess("cesi", true)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("stop status:", r)
}

func Test_StopGroupProcess(t *testing.T) {
	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r, err := client.StopProcessGourp("cesi", true)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range r {
		t.Log("processname:", v.ProcessName)
		t.Log("group:", v.Group)
		t.Log("description:", v.Description)
		t.Log("status:", v.Status)
	}
}

func Test_StopAllProcesses(t *testing.T) {
	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r, err := client.StopAllProcesses(true)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range r {
		t.Log("processname:", v.ProcessName)
		t.Log("group:", v.Group)
		t.Log("description:", v.Description)
		t.Log("status:", v.Status)
	}
}

func Test_ReloadConfig(t *testing.T) {
	client, err := supervisor.NewSupervisor("192.168.250.178:9001")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	r, err := client.ReloadConfig()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("reload config status:", r)
}