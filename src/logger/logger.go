package logger

import (
    "fmt"
    "log"
    "os"
    "path"
    "runtime"
)

const (
    LOG_LEVEL_FATAL   = 0x0002
    LOG_LEVEL_WARNING = 0x0010
    LOG_LEVEL_TRACE   = 0x0080
    LOG_LEVEL_NOTICE  = 0x0100
    LOG_LEVEL_DEBUG   = 0x0800
    LOG_LEVEL_VERBOSE = 0x1000
)

var (
    logLevel      = LOG_LEVEL_VERBOSE
    logStringChan = make(chan *LoggerString, 4096)
    logChangeChan = make(chan *LoggerChange, 16)
    logExitChan   = make(chan bool, 2)
    logWaitChan   = make(chan bool, 1)

    logLoggers = map[int]*log.Logger{
        LOG_LEVEL_FATAL:   log.New(os.Stdout, "FATAL:", log.LstdFlags),
        LOG_LEVEL_WARNING: log.New(os.Stdout, "WARNING:", log.LstdFlags),
        LOG_LEVEL_TRACE:   log.New(os.Stdout, "TRACE:", log.LstdFlags),
        LOG_LEVEL_NOTICE:  log.New(os.Stdout, "NOTICE:", log.LstdFlags),
        LOG_LEVEL_DEBUG:   log.New(os.Stdout, "DEBUG:", log.LstdFlags),
        LOG_LEVEL_VERBOSE: log.New(os.Stdout, "VERBOSE:", log.LstdFlags),
    }
)

type LoggerString struct {
    file  string
    line  int
    level int
    fmt   string
    args  []interface{}
}

type LoggerChange struct {
    level  int
    logger *log.Logger
}

func init() {
    go loggerLoop()
}

func SetLogLevel(level int) {
    logLevel = level
}

func logf(level int, f string, args ...interface{}) {
    if logLevel >= level {
        _, file, line, ok := runtime.Caller(2)

        if ok {
            file = path.Base(file)
        } else {
            file = "???"
            line = 1
        }

        logStringChan <- &LoggerString{
            file:  file,
            line:  line,
            level: level,
            fmt:   f,
            args:  args,
        }
    }
}

// Logs notice level log
func Logf(f string, args ...interface{}) {
    logf(LOG_LEVEL_NOTICE, f, args...)
}

func LogDebugf(f string, args ...interface{}) {
    logf(LOG_LEVEL_DEBUG, f, args...)
}

func LogTracef(f string, args ...interface{}) {
    logf(LOG_LEVEL_TRACE, f, args...)
}

func LogNoticef(f string, args ...interface{}) {
    logf(LOG_LEVEL_NOTICE, f, args...)
}

func LogWarningf(f string, args ...interface{}) {
    logf(LOG_LEVEL_WARNING, f, args...)
}

func LogFatalf(f string, args ...interface{}) {
    logf(LOG_LEVEL_FATAL, f, args...)
}

// Logs package bytes in debug log
func LogPackagef(buf []byte, f string, args ...interface{}) {
    fmt := fmt.Sprintf("%v[len:%v] % x", f, len(buf), buf)
    logf(LOG_LEVEL_VERBOSE, fmt, args...)
}

// Flushes all cached logs.
func LogClose() {
    logExitChan <- true
    <-logWaitChan
}

func logPrintLoggerString(str *LoggerString) {
    if logger, ok := logLoggers[str.level]; ok {
        fmt := fmt.Sprintf("[%v:%v] %v", str.file, str.line, str.fmt)
        logger.Printf(fmt, str.args...)
    }
}

func loggerLoop() {
LoggingLoop:
    for {
        select {
        case str := <-logStringChan:
            logPrintLoggerString(str)

        case change := <-logChangeChan:
            logLoggers[change.level] = change.logger

        case <-logExitChan:
            break LoggingLoop
        }
    }

    // flush all log messages
FlushLoop:
    for {
        select {
        case str := <-logStringChan:
            logPrintLoggerString(str)

        default:
            break FlushLoop
        }
    }

    logWaitChan <- true
}
