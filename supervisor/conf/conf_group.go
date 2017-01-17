/*
@author: foolbread
@time: 2017-01-17
@file: supervisor-api/supervisor/conf/conf_group.go
*/
package conf

import (
	"strings"
)

var groupKeys []string = []string{"programs","priority"}

type confGroup struct {
	name string
	content map[string]string
}

func newConfGroup(name string)*confGroup{
	r := new(confGroup)
	r.name = name
	r.content = make(map[string]string)
	for _,v := range groupKeys{
		r.content[v] = ""
	}

	return r
}

func (c *confGroup)encode()string{
	var strs []string
	strs = append(strs,"[group:"+c.name+"]")
	for _, v :=range groupKeys{
		if len(c.content[v]) > 0{
			strs = append(strs, v+"="+c.content[v])
		}
	}

	return strings.Join(strs,"\n")
}
