package logging

import (
	"log"
	"os"
)

// From http://www.goinggo.net/2013/11/using-log-package-in-go.html

func InitLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

var (
	Trace = InitLogger("TRACE: ")
	Info  = InitLogger("INFO : ")
	Error = InitLogger("ERROR: ")
)
