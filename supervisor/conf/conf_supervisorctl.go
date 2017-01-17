/*
@author: foolbread
@time: 2017-01-17
@file: supervisor-api/supervisor/conf/conf_supervisorctl.go
*/
package conf

import (
	"strings"
)

var supervisorctlKeys []string = []string{"serverurl","username","password","prompt"}

type confSupervisorctl struct {
	content map[string]string
}

func newConfSupervisorctl()*confSupervisorctl{
	r := new(confSupervisorctl)
	r.content = make(map[string]string)
	for _, v := range supervisorctlKeys{
		r.content[v] = ""
	}

	return r
}

func (c *confSupervisorctl)encode()string{
	var strs []string
	strs = append(strs, "[supervisorctl]")

	for _, v := range supervisorctlKeys{
		if len(c.content[v]) > 0{
			strs = append(strs, v+"="+c.content[v])
		}
	}

	return strings.Join(strs, "\n")
}