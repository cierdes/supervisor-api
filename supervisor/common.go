/*
@author: foolbread
@time: 2017-01-08
@file: supervisor-api/supervisor/common.go
*/
package supervisor

const (
	STOPPED_CODE = 0    //The process has been stopped due to a stop request or has never been started.
	STARTING_CODE = 10  //The process is starting due to a start request.
	RUNNING_CODE = 20   //The process is running.
	BACKOFF_CODE = 30   //The process entered the STARTING state but subsequently exited too quickly to move to the RUNNING state.
	STOPPING_CODE = 40  //The process is stopping due to a stop request.
	EXITED_CODE = 100   //The process exited from the RUNNING state (expectedly or unexpectedly).
	FATAL_CODE = 200    //The process could not be started successfully.
	UNKNOWN_CODE = 1000 //The process is in an unknown state (supervisord programming error).
)

const (
	STOPPED_NAME = "STOPPED"      //The process has been stopped due to a stop request or has never been started.
	STARTING_NAME = "STARTING"    //The process is starting due to a start request.
	RUNNING_NAME = "RUNNING"      //The process is running.
	BACKOFF_NAME = "BACKOFF"      //The process entered the STARTING state but subsequently exited too quickly to move to the RUNNING state.
	STOPPING_NAME = "STOPPING"    //The process is stopping due to a stop request.
	EXITED_NAME = "EXITED"        //The process exited from the RUNNING state (expectedly or unexpectedly).
	FATAL_NAME = "FATAL"          //The process could not be started successfully.
	UNKNOWN_NAME = "UNKNOWN"      //The process is in an unknown state (supervisord programming error).
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SupervisorProcessInfo struct {
	ProcessName string	`xmlrpc:"name"`
	Group string		`xmlrpc:"group"`
	Description string 	`xmlrpc:"description"`
	Start int		`xmlrpc:"start"`
	Stop int		`xmlrpc:"stop"`
	Now int			`xmlrpc:"now"`
	State int		`xmlrpc:"state"`
	StateName string	`xmlrpc:"statename"`
	Spawnerr string		`xmlrpc:"spawnerr"`
	ExitStatus int		`xmlrpc:"exitstatus"`
	LogFile string		`xmlrpc:"logfile"`
	StdoutLogFile string	`xmlrpc:"stdout_logfile"`
	StderrLogFile string	`xmlrpc:"stderr_logfile"`
	Pid int			`xmlrpc:"pid"`
}

type SupervisorProcessStatus struct {
	ProcessName string 	`xmlrpc:"name"`
	Group string 		`xmlrpc:"group"`
	Description string 	`xmlrpc:"description"`
	Status int		`xmlrpc:"status"`
}