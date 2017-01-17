/*
@author: foolbread
@time: 2017-01-17
@file: supervisor-api/supervisor/conf/conf_unixhttpserver.go
*/
package conf

import (
	"strings"
)

var unxitHttpServerKeys []string = []string{"file", "chmod", "chown", "username", "password"}


type confUnixHttpServer struct {
	content map[string]string
}

func newConfUnixHttpServer()*confUnixHttpServer{
	r := new(confUnixHttpServer)
	r.content = make(map[string]string)
	for _, v := range unxitHttpServerKeys {
		r.content[v] = ""
	}

	return r
}

func (c *confUnixHttpServer)encode()string{
	var strs []string
	strs = append(strs,"[unix_http_server]")

	for _, v :=range unxitHttpServerKeys{
		if len(c.content[v]) > 0{
			strs = append(strs, v+"="+c.content[v])
		}
	}

	return strings.Join(strs,"\n")
}