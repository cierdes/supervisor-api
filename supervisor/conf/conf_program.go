/*
@author: foolbread
@time: 2017-01-17
@file: supervisor-api/supervisor/conf/conf_program.go
*/
package conf

import (
	"strings"
)

var programKeys []string = []string{"command","process_name","numprocs","directory","umask","priority","autostart",
	                            "autorestart","startsecs","startretries","exitcodes","stopsignal","stopwaitsecs",
                                    "stopasgroup","killasgroup","user","redirect_stderr","stdout_logfile","stdout_logfile_maxbytes",
	                            "stdout_logfile_backups","stdout_capture_maxbytes","stdout_events_enabled","stderr_logfile",
	                            "stderr_logfile_maxbytes","stderr_logfile_backups","stderr_capture_maxbytes","stderr_events_enabled",
                                    "environment","serverurl"}

type confProgram struct {
	name string
	content map[string]string
}

func newConfProgram(name string)*confProgram{
	r := new(confProgram)
	r.name = name
	r.content = make(map[string]string)
	for _,v := range programKeys{
		r.content[v] = ""
	}

	return r
}

func (c *confProgram)encode()string{
	var strs []string
	strs = append(strs,"[program:"+c.name+"]")
	for _,v :=range programKeys{
		if len(c.content[v]) > 0{
			strs = append(strs, v+"="+c.content[v])
		}
	}

	return strings.Join(strs,"\n")
}