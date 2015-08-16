package logs

import (
	"log"
	"os"
)

// From http://www.goinggo.net/2013/11/using-log-package-in-go.html

func initLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

var (
	Trace = initLogger("TRACE: ")
	Info  = initLogger("INFO : ")
	Error = initLogger("ERROR: ")
	Warn  = initLogger("WARN : ")
)
