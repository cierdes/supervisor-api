/*
@author: foolbread
@time: 2017-01-17
@file: supervisor-api/supervisor/conf/conf_eventlistener.go
*/
package conf

import (
	"strings"
)

var eventlistenerKeys []string = []string{"command","process_name","numprocs","events","buffer_size","directory","umask","priority","autostart",
	 				  "autorestart","startsecs","startretries","exitcodes","stopsignal","stopwaitsecs",
				          "stopasgroup","killasgroup","user","redirect_stderr","stdout_logfile","stdout_logfile_maxbytes",
				          "stdout_logfile_backups","stdout_events_enabled","stderr_logfile",
				          "stderr_logfile_maxbytes","stderr_logfile_backups","stderr_events_enabled", "environment","serverurl"}

type confEventListener struct {
	name string
	content map[string]string
}

func newConfEventListener(name string)*confEventListener{
	r := new(confEventListener)
	r.name = name
	r.content = make(map[string]string)
	for _,v := range eventlistenerKeys{
		r.content[v] = ""
	}

	return r
}

func (c *confEventListener)encode()string{
	var strs []string
	strs = append(strs, "[eventlistener:"+c.name+"]")
	for _,v :=range eventlistenerKeys{
		if len(c.content[v]) > 0{
			strs = append(strs, v+"="+c.content[v])
		}
	}

	return strings.Join(strs,"\n")
}
