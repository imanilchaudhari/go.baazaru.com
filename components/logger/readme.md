# Golang file logger

## Settings

```go
// Set directory path to save log files
logger.Path("./log")

// Set log level by constant
logger.Level(logger.DEBUG)

// Set log level by string
logger.LevelAsString("debug")

// Set function to prepare format log line
logger.Format(func(level int, line string, message string) string{
    return fmt.Sprintf("Only message: %s", message)
})

// Set log file size limit
logger.SizeLimit(2 * logger.MB)

// Set log output to stdout 
logger.Stdout(true)
```

## Usage
```go
// Write debug line
// Usually use to log debug data
logger.Debug("debug line")

// Write information
// Usually use to log some state
logger.Info("info line")

// Write warning
// Usually use to log warnings about unexpected application state (ex.: brudforce, incorrect request, bad loging&password authorization) 
logger.Warn("warn line")

// Write error
// Use only in a case of a return error what doesn't effect application run
logger.Error("error line")

// Write fatal error
// Use only in a case of a return error what do effect application run
logger.Fatal("fatal line")
```

## List all methods

* logger.Path("./log") - sets directory path to save log files
* logger.Level(logger.DEBUG) - sets log level by constant
* logger.LevelAsString("debug") - sets log level by string
* logger.Syslog("tag") - writes to localhost syslog with tag
* logger.Format(func(level int, line string, message string) string) - sets function to prepare format log line
* logger.SizeLimit(2 * logger.MB) - sets log file size limit (if value is less zero then size limit is off)
* logger.Stdout(true) - sets log output to stdout
* logger.TTL(3600) - sets time-to-live of log files (all old files will be removed)
* logger.Extension("txt") = sets another log files extension (.log is default)
* logger.Debug("debug line") - writes message with debug data
* logger.DebugFmt("debug line %d", 1) - writes message with debug data
* logger.Info("info line") - writes message with information about state or similar
* logger.InfoFmt("info line %d", 1) - writes message with information about state or similar
* logger.Warn("warn line") - usually use to write warning message about unexpected application state (ex.: brudforce, incorrect request, bad loging&password authorization) 
* logger.WarnFmt("warn line %d", 1) - usually use to write warning message about unexpected application state (ex.: brudforce, incorrect request, bad loging&password authorization) 
* logger.Error("error line") - use only in a case of a return error what doesn't effect application run
* logger.ErrorFmt("error line %d", 1) - use only in a case of a return error what doesn't effect application run
* logger.Fatal("fatal line") - use only in a case of a return error what do effect application run
* logger.FatalFmt("fatal line %d", 1) - use only in a case of a return error what do effect application run
* logger.Flush() - finish all process
* logger.Notifier(func(message string){ /* do something */ }, "warn") - sets callback on all message with level more or equal "warn" ("warn", "error", "fatal")

SOURCE : https://github.com/leprosus/golang-log