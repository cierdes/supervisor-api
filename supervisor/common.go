/*
@author: foolbread
@time: 2017-01-08
@file: supervisor-api/supervisor/common.go
*/
package supervisor

const (
	PROCESS_STOPPED_CODE = 0    //The process has been stopped due to a stop request or has never been started.
	PROCESS_STARTING_CODE = 10  //The process is starting due to a start request.
	PROCESS_RUNNING_CODE = 20   //The process is running.
	PROCESS_BACKOFF_CODE = 30   //The process entered the STARTING state but subsequently exited too quickly to move to the RUNNING state.
	PROCESS_STOPPING_CODE = 40  //The process is stopping due to a stop request.
	PROCESS_EXITED_CODE = 100   //The process exited from the RUNNING state (expectedly or unexpectedly).
	PROCESS_FATAL_CODE = 200    //The process could not be started successfully.
	PROCESS_UNKNOWN_CODE = 1000 //The process is in an unknown state (supervisord programming error).
)

const (
	PROCESS_STOPPED_NAME = "STOPPED"      //The process has been stopped due to a stop request or has never been started.
	PROCESS_STARTING_NAME = "STARTING"    //The process is starting due to a start request.
	PROCESS_RUNNING_NAME = "RUNNING"      //The process is running.
	PROCESS_BACKOFF_NAME = "BACKOFF"      //The process entered the STARTING state but subsequently exited too quickly to move to the RUNNING state.
	PROCESS_STOPPING_NAME = "STOPPING"    //The process is stopping due to a stop request.
	PROCESS_EXITED_NAME = "EXITED"        //The process exited from the RUNNING state (expectedly or unexpectedly).
	PROCESS_FATAL_NAME = "FATAL"          //The process could not be started successfully.
	PROCESS_UNKNOWN_NAME = "UNKNOWN"      //The process is in an unknown state (supervisord programming error).
)

const (
	SUPERVISOR_SHUTDOWN_CODE = -1	//Supervisor is in the process of shutting down.
	SUPERVISOR_RESTARTING_CODE = 0	//Supervisor is in the process of restarting.
	SUPERVISOR_RUNNING_CODE = 1	//Supervisor is working normally.
	SUPERVISOR_FATAL_CODE = 2	//Supervisor has experienced a serious error.
)

const (
	SUPERVISOR_SHUTDOWN_NAME = "SHUTDOWN"		//Supervisor is in the process of shutting down.
	SUPERVISOR_RESTARTING_NAME = "RESTARTING"	//Supervisor is in the process of restarting.
	SUPERVISOR_RUNNING_NAME = "RUNNING"		//Supervisor is working normally.
	SUPERVISOR_FATAL_NAME = "FATAL"			//Supervisor has experienced a serious error.
)
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SupervisorStatus struct {
	StateCode int		`xmlrpc:"statecode"`
	StateName string	`xmlrpc:"statename"`
}

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