package misc

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

type Logger struct {
	Trace *log.Logger
	Debug *log.Logger
	Error *log.Logger
}

var logger Logger

func InitLog() {
	config := ReadConfig("")
	logger = Logger{}
	logPath := path.Join(config.LogPath, "log.txt")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	traceW := io.Writer(os.Stdout)
	debugW := io.Writer(os.Stdout)
	errorW := io.Writer(os.Stdout)
	if err != nil {
		log.Println("Failed to open log file")
	} else {
		traceW = io.Writer(file)
		debugW = io.MultiWriter(file, os.Stdout)
		errorW = io.Writer(file)
	}

	logger.Trace = log.New(traceW,
		"[TRACE]: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.Debug = log.New(debugW,
		"[DEBUG]: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.Error = log.New(errorW,
		"[ERROR]: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func LogError(err error) {
	if logger.Error == nil {
		fmt.Println("Error logger was not initialized")
		return
	}
	logger.Error.Println(err)
}

func LogTrace(info string) {
	if logger.Trace == nil {
		fmt.Println("Trace logger was not initialized")
		return
	}
	logger.Trace.Println(info)
}

func LogDebug(info interface{}) {
	if logger.Debug == nil {
		fmt.Println("Debug logger was not initialized")
		return
	}
	logger.Debug.Println(info)
}
