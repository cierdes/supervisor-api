/*
@author: foolbread
@time: 2017-01-17
@file: supervisor-api/supervisor/conf/conf_inethttpserver.go
*/
package conf

import "strings"

var inetHttpServerKeys []string = []string{"port","username","password"}

type confInetHttpServer struct {
	content map[string]string
}

func newConfInetHttpServer()*confInetHttpServer{
	r := new(confInetHttpServer)
	r.content = make(map[string]string)
	for _, v := range inetHttpServerKeys{
		r.content[v] = ""
	}

	return r
}

func (c *confInetHttpServer)encode()string{
	var strs []string
	strs = append(strs, "[inet_http_server]")

	strs = append(strs, encode(c.content)...)

	return strings.Join(strs,"\n")
}