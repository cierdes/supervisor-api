/*
@author: foolbread
@time: 2017-01-17
@file: supervisor-api/supervisor/conf/conf_rpcinterface.go
*/
package conf

import (
	"strings"
)

var rpcinterfaceKeys []string = []string{"supervisor.rpcinterface_factory","retries"}

type confRPCInterface struct {
	name string
	content map[string]string
}

func newConfRPCInterface(name string)*confRPCInterface{
	r := new(confRPCInterface)
	r.name = name
	r.content = make(map[string]string)
	for _,v := range rpcinterfaceKeys{
		r.content[v] = ""
	}

	return r
}

func (c *confRPCInterface)encode()string{
	var strs []string
	strs = append(strs,"[rpcinterface:"+c.name+"]")
	for _,v := range rpcinterfaceKeys{
		if len(c.content[v]) > 0{
			strs = append(strs, v+"="+c.content[v])
		}
	}

	return strings.Join(strs,"\n")
}