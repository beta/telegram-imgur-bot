package apis

import (
	"fmt"
	"log"
	"os"
)

var (
	stdoutLogger = log.New(os.Stdout, "", log.LstdFlags)
	stderrLogger = log.New(os.Stderr, "", log.LstdFlags)
)

// LogInfo outputs a piece of log at the level of info.
func (api *API) LogInfo(req Request, args ...interface{}) {
	stdoutLogger.Print(append([]interface{}{fmt.Sprintf("INFO [%s]", req.ReqID())}, args...)...)
}

// LogWarn outputs a piece of log at the level of warning.
func (api *API) LogWarn(req Request, args ...interface{}) {
	stdoutLogger.Print(append([]interface{}{fmt.Sprintf("WARN [%s]", req.ReqID())}, args...)...)
}

// LogError outputs a piece of log at the level of error.
func (api *API) LogError(req Request, args ...interface{}) {
	stderrLogger.Print(append([]interface{}{fmt.Sprintf("ERRO [%s]", req.ReqID())}, args...)...)
}

// LogInfof outputs a piece of log for a message at the level of info.
func (api *API) LogInfof(req Request, format string, args ...interface{}) {
	stdoutLogger.Printf("INFO [%s] %s", req.ReqID(), fmt.Sprintf(format, args...))
}

// LogWarnf outputs a piece of log for a message at the level of warning.
func (api *API) LogWarnf(req Request, format string, args ...interface{}) {
	stdoutLogger.Printf("WARN [%s] %s", req.ReqID(), fmt.Sprintf(format, args...))
}

// LogErrorf outputs a piece of log for a message at the level of error.
func (api *API) LogErrorf(req Request, format string, args ...interface{}) {
	stderrLogger.Printf("ERRO [%s] %s", req.ReqID(), fmt.Sprintf(format, args...))
}
