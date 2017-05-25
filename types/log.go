/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "sync"

type Log_t struct {
    *File_t

     sync.Mutex

     level  int
     path   string

     Log
}

const (
    STDERR = iota
    INFO
    WARN
    ERROR
    DEBUG
    PANIC
    FATAL
)

var Levels = [...]string{
    STDERR : "stderr",
    INFO   : "info",
    WARN   : "warn",
    ERROR  : "error",
    DEBUG  : "debug",
    PANIC  : "panic",
    FATAL  : "fatal",
}

func NewLog() *Log_t {
    return &Log_t{
        File_t:  NewFile(nil),
        level : INFO,
    }
}

func (l *Log_t) SetLevel(level int) int {
    if level > FATAL {
        return Error
    }

    l.level = level

    return Ok
}

func (l *Log_t) GetLevel() int {
    return l.level
}

func (l *Log_t) SetPath(path string) int {
    if path == "" {
        return Error
    }

    l.path = path

    return Ok
}

func (l *Log_t) GetPath() string {
    return l.path
}

func (l *Log_t) Set(log Log) int {
    if log == nil {
        return Error
    }

    l.Log = log

    return Ok
}

func (l *Log_t) Get() Log {
    return l.Log
}

func (l *Log_t) print(format string, p ...interface{}) int {
    if format == "" || p == nil {
        return Error
    }

    if log := l.Get(); log != nil {
        if log.Dump() == Error {
            return Error
        }

        return Ok
    }

    // default method
    if l.Dump() == Error {
        return Error
    }

    return Ok
}

func (l *Log_t) Stderr(format string, d ...interface{}) {
    return
}

func (l *Log_t) Info(format string, i ...interface{}) {
				l.print(format, i)
}

func (l *Log_t) Warn(format string, w ...interface{}) {
    if l.level < WARN {
        return
    }

    if l.print(format, w) == Error {
        return
    }

    return
}

func (l *Log_t) Error(format string, e ...interface{}) {
    l.print(format, e)
}

func (l *Log_t) Debug(format string, d ...interface{}) {
    if l.level < DEBUG {
        return
    }

    if l.print(format, d) == Error {
        return
    }

    return
}

func (l *Log_t) Panic(format string, d ...interface{}) {
    if l.print(format, d) == Error {
        return
    }

    return
}

func (l *Log_t) Fatal(format string, d ...interface{}) {
    if l.print(format, d) == Error {
        return
    }

    return
}

func (l *Log_t) Dump() int {
    //log.
    return Ok
}
