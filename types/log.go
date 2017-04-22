/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "sync"

type Log struct {
    *File

     sync.Mutex

     level  int
     path   string
     log    LogIf
}

type LogIf interface {
    Dump() int
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

func NewLog() *Log {
    return &Log{
        File:  NewFile(nil),
        level: INFO,
    }
}

func (l *Log) SetLevel(level int) int {
    if level > FATAL {
        return Error
    }

    l.level = level

    return Ok
}

func (l *Log) GetLevel() int {
    return l.level
}

func (l *Log) SetPath(path string) int {
    if path == "" {
        return Error
    }

    l.path = path

    return Ok
}

func (l *Log) GetPath() string {
    return l.path
}

func (l *Log) Set(log LogIf) int {
    if log == nil {
        return Error
    }

    l.log = log

    return Ok
}

func (l *Log) Get() LogIf {
    return l.log
}

func (l *Log) print(format string, p ...interface{}) int {
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

func (l *Log) Stderr(format string, d ...interface{}) {
    return
}

func (l *Log) Info(format string, i ...interface{}) {
				l.print(format, i)
}

func (l *Log) Warn(format string, w ...interface{}) {
    if l.level < WARN {
        return
    }

    if l.print(format, w) == Error {
        return
    }

    return
}

func (l *Log) Error(format string, e ...interface{}) {
    l.print(format, e)
}

func (l *Log) Debug(format string, d ...interface{}) {
    if l.level < DEBUG {
        return
    }

    if l.print(format, d) == Error {
        return
    }

    return
}

func (l *Log) Panic(format string, d ...interface{}) {
    if l.print(format, d) == Error {
        return
    }

    return
}

func (l *Log) Fatal(format string, d ...interface{}) {
    if l.print(format, d) == Error {
        return
    }

    return
}

func (l *Log) Dump() int {
    //log.
    return Ok
}
