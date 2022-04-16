package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/akromibn37/hospitalityCollaboration/pkg/file"
)

type Level int

var (
	FErr, FInf, FPerf  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logInterface, logError, logPerformance *log.Logger

	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup initialize the log instance
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := "error_" + getLogFileName()
	FErr, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logError = log.New(FErr, DefaultPrefix, log.LstdFlags)

	fileName = "interface_" + getLogFileName()
	FInf, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logInterface = log.New(FInf, DefaultPrefix, log.LstdFlags)

	fileName = "performance_" + getLogFileName()
	FPerf, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logPerformance = log.New(FPerf, DefaultPrefix, log.LstdFlags)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	// logger.Println(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	prefix := setPrefix(INFO)
	logInterface.SetPrefix(prefix)
	logInterface.Println(v)
}

// Info output logs performance at info level
func PerfInfo(v ...interface{}) {
	prefix := setPrefix(INFO)
	logPerformance.SetPrefix(prefix)
	logPerformance.Println(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logError.Println(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	prefix := setPrefix(ERROR)
	logError.SetPrefix(prefix)
	logError.Println(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	// logger.Fatalln(v)
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) string {
	if level == 3 {
		_, file, line, ok := runtime.Caller(DefaultCallerDepth)
		if ok {
			logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
		} else {
			logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
		}
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	return logPrefix

}
