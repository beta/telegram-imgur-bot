package apis

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/tucnak/telebot.v2"
)

var (
	stdoutLogger = log.New(os.Stdout, "", log.LstdFlags)
	stderrLogger = log.New(os.Stderr, "", log.LstdFlags)
)

// LogInfo outputs a piece of log at the level of info.
func (api *API) LogInfo(m *telebot.Message, args ...interface{}) {
	stdoutLogger.Print(append([]interface{}{fmt.Sprintf("INFO [%d]", m.ID)}, args...)...)
}

// LogWarn outputs a piece of log at the level of warning.
func (api *API) LogWarn(m *telebot.Message, args ...interface{}) {
	stdoutLogger.Print(append([]interface{}{fmt.Sprintf("WARN [%d]", m.ID)}, args...)...)
}

// LogError outputs a piece of log at the level of error.
func (api *API) LogError(m *telebot.Message, args ...interface{}) {
	stderrLogger.Print(append([]interface{}{fmt.Sprintf("ERRO [%d]", m.ID)}, args...)...)
}

// LogInfof outputs a piece of log for a message at the level of info.
func (api *API) LogInfof(m *telebot.Message, format string, args ...interface{}) {
	stdoutLogger.Printf("INFO [%d] %s", m.ID, fmt.Sprintf(format, args...))
}

// LogWarnf outputs a piece of log for a message at the level of warning.
func (api *API) LogWarnf(m *telebot.Message, format string, args ...interface{}) {
	stdoutLogger.Printf("WARN [%d] %s", m.ID, fmt.Sprintf(format, args...))
}

// LogErrorf outputs a piece of log for a message at the level of error.
func (api *API) LogErrorf(m *telebot.Message, format string, args ...interface{}) {
	stderrLogger.Printf("ERRO [%d] %s", m.ID, fmt.Sprintf(format, args...))
}
