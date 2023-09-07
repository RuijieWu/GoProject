// 用于隔离的配置文件，进行破解时只需要动config.go
package main

import "sync"

var (
	source_string string         = "s878926199?" //案例
	target_string string         = "0e545993274517709034328855841020"
	ans                          = make(chan string)
	wg            sync.WaitGroup = sync.WaitGroup{}
	num           int            = 0
	count         int            = 0
	sum           int            = 0
	flag          int            = 0
)
