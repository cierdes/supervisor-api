/*
@author: foolbread
@time: 2017-01-17
@file: supervisor-api/supervisor/conf/conf_include.go
*/
package conf

import (
	"strings"
)

var includeKeys []string = []string{"files"}

type confInclude struct {
	content map[string]string
}

func newConfInclude()*confInclude{
	r := new(confInclude)
	r.content = make(map[string]string)
	for _, v := range includeKeys{
		r.content[v] = ""
	}

	return r
}

func (c *confInclude)encode()string{
	var strs []string
	strs = append(strs, "[include]")
	for _, v := range includeKeys{
		if len(c.content[v]) > 0{
			strs = append(strs, v+"="+c.content[v])
		}
	}

	return strings.Join(strs, "\n")
}