/*
@author: foolbread
@time: 2017-01-17
@file: supervisor-api/supervisor/conf/conf_supervisord.go
*/
package conf

import (
	"strings"
)

var supervisordKeys []string = []string{"logfile","logfile_maxbytes","logfile_backups","loglevel","pidfile",
					"nodaemon","minfds","minprocs","umask","user","identifier","directory",
					"nocleanup","childlogdir","strip_ansi","environment"}

type confSupervisord struct {
	content map[string]string
}

func newConfSupervisord()*confSupervisord{
	r := new(confSupervisord)
	r.content = make(map[string]string)
	for _,v :=range supervisordKeys{
		r.content[v] = ""
	}

	return r
}

func (c *confSupervisord)encode()string{
	var strs []string
	strs = append(strs, "[supervisord]")

	for _,v :=range supervisordKeys{
		if len(c.content[v]) > 0{
			strs = append(strs, v+"="+c.content[v])
		}
	}

	return strings.Join(strs, "\n")
}