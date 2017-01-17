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