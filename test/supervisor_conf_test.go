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