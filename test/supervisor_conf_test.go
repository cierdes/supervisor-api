/*
@author: foolbread
@time: 2017-01-17
@file: supervisor-api/test/supervisor_conf_test.go
*/
package test

import (
	"github.com/cierdes/supervisor-api/supervisor/conf"

	"testing"
)

func Test_UnixHttpServer(t *testing.T){
	var conf conf.SupervisorConf

	err := conf.WriteUnixHttpServer("file","/tmp/supervisor.sock")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteUnixHttpServer("chmod","0777")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteUnixHttpServer("chown","nobody:nogroup")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteUnixHttpServer("username","user")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteUnixHttpServer("password","123")
	if err != nil{
		t.Fatal(err)
	}

	t.Log(conf.EncodeToString())
}

func Test_InetHttpServer(t *testing.T){
	var conf conf.SupervisorConf

	err := conf.WriteInetHttpServer("port","127.0.0.1:9001")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteInetHttpServer("username","user")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteInetHttpServer("password","123")
	if err != nil{
		t.Fatal(err)
	}

	t.Log(conf.EncodeToString())
}

func Test_Supervisord(t *testing.T){
	var conf conf.SupervisorConf

	err := conf.WriteSupervisord("logfile","/tmp/supervisord.log")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("logfile_maxbytes","50MB")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("logfile_backups","10")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("loglevel","info")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("pidfile","/tmp/supervisord.pid")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("nodaemon","false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("minfds","1024")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("minprocs","200")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("umask","022")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("user","chrism")
	if err != nil{
		t.Fatal(err)
	}

	err =conf.WriteSupervisord("identifier","supervisor")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("directory","/tmp")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("nocleanup","true")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("childlogdir","/tmp")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("strip_ansi","false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisord("environment",`KEY1="value1",KEY2="value2"`)
	if err != nil{
		t.Fatal(err)
	}

	t.Log(conf.EncodeToString())
}

func Test_Supervisorctl(t *testing.T){
	var conf conf.SupervisorConf

	err := conf.WriteSupervisorctl("serverurl","unix:///tmp/supervisor.sock")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisorctl("username","chris")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisorctl("password","123")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteSupervisorctl("prompt","mysupervisor")
	if err != nil{
		t.Fatal(err)
	}

	t.Log(conf.EncodeToString())
}

func Test_Program(t *testing.T){
	var conf conf.SupervisorConf

	err := conf.WriteProgram("hello","command","/bin/cat")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","process_name","%(program_name)s")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","numprocs","1")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "directory", "/tmp")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "umask","022")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","priority", "999")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","autostart","true")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "autorestart", "unexpected")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "startsecs", "10")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "startretries", "3")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "exitcodes", "0,2")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "stopsignal", "TERM")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "stopwaitsecs", "10")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "stopasgroup", "false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","killasgroup","false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","user", "chrism")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "redirect_stderr","false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","stdout_logfile","/a/path")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","stdout_logfile_maxbytes","1MB")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","stdout_logfile_backups","10")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "stdout_capture_maxbytes","1MB")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","stdout_events_enabled", "false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "stderr_logfile","/a/path")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "stderr_logfile_maxbytes", "1MB")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "stderr_logfile_backups","10")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","stderr_logfile_maxbytes","1MB")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello","stderr_capture_maxbytes", "1MB")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "stderr_events_enabled", "false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "environment", `A="1",B="2"`)
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("hello", "serverurl", "AUTO")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteProgram("kkkk","serverurl","AUTO")
	if err != nil{
		t.Fatal(err)
	}


	t.Log(conf.EncodeToString())
}

func Test_Include(t *testing.T){
	var conf conf.SupervisorConf

	err := conf.WriteInclude("files","/an/absolute/filename.conf /an/absolute/8.conf foo.cnf config??.conf")
	if err != nil{
		t.Fatal(err)
	}

	t.Log(conf.EncodeToString())
}

func Test_Group(t *testing.T){
	var conf conf.SupervisorConf

	err := conf.WriteGroup("hello","programs","bar,baz")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteGroup("hello","priority","999")
	if err != nil{
		t.Fatal(err)
	}

	t.Log(conf.EncodeToString())
}

func Test_EventListener(t *testing.T){
	var conf conf.SupervisorConf

	err := conf.WriteEventListener("hello","command","/bin/cat")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","process_name","%(program_name)s")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","numprocs","1")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","events","PROCESS_STATE")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","buffer_size","10")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "directory", "/tmp")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "umask","022")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","priority", "999")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","autostart","true")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "autorestart", "unexpected")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "startsecs", "10")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "startretries", "3")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "exitcodes", "0,2")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "stopsignal", "TERM")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "stopwaitsecs", "10")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "stopasgroup", "false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","killasgroup","false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","user", "chrism")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "redirect_stderr","false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","stdout_logfile","/a/path")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","stdout_logfile_maxbytes","1MB")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","stdout_logfile_backups","10")
	if err != nil{
		t.Fatal(err)
	}


	err = conf.WriteEventListener("hello","stdout_events_enabled", "false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "stderr_logfile","/a/path")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "stderr_logfile_maxbytes", "1MB")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "stderr_logfile_backups","10")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello","stderr_logfile_maxbytes","1MB")
	if err != nil{
		t.Fatal(err)
	}


	err = conf.WriteEventListener("hello", "stderr_events_enabled", "false")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "environment", `A="1",B="2"`)
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("hello", "serverurl", "AUTO")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteEventListener("kkkk","serverurl","AUTO")
	if err != nil{
		t.Fatal(err)
	}


	t.Log(conf.EncodeToString())
}

func Test_RPCInterface(t *testing.T){
	var conf conf.SupervisorConf

	err := conf.WriteRPCInterface("hello","supervisor.rpcinterface_factory","my.package:make_another_rpcinterface")
	if err != nil{
		t.Fatal(err)
	}

	err = conf.WriteRPCInterface("hello","retries","1")
	if err != nil{
		t.Fatal(err)
	}

	t.Log(conf.EncodeToString())
}