/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "os"

type Module_t struct {
    Index     int          //pointer module position
    Cursor    int          //pointer context position
    Context   Context
    Commands  []Command_t
    Type      int64
}

var Modules []Module

const (
    INIT = iota
    MAIN
    EXIT
)

var Sentinel = [...]bool{
    INIT: false,
    MAIN: false,
    EXIT: false,
}

func (r *Module_t) Init(o *Option_t) int {
    if !Sentinel[INIT] {
        return Ok
    }

    Sentinel[INIT] = false

    modules := GetModules(Modules, SYSTEM_MODULE)
    if modules == nil {
        return Error
    }

    for _, v := range modules {
        if self := v.Self(); self != nil {
            if v.Init(o) == Error {
                os.Exit(0)
            }
        }
    }

    var configure *Configure_t
    if configure = o.Configure_t; configure == nil {
        configure = NewConfigure(o.Log_t)
    }

    for _, v := range modules {
        if self := v.Self(); self != nil {
            go v.Main(configure)
        }
    }

    select {

    case e := <- configure.Event:
        if op := e.GetOpcode(); op != LOAD {
            return Ignore
        }
    }

    if Block(configure, Modules, r.Type, CONFIG_BLOCK) == Error {
        return Error
    }

    // NOTE
    configure.value = o

    return Ok
}

func (r *Module_t) Main(c *Configure_t) int {
    if !Sentinel[MAIN] {
        return Ok
    }

    Sentinel[MAIN] = false

    o := c.value.(*Option_t)
    if o == nil {
        // TODO add log
        return Error
    }

    modules := GetModules(Modules, SYSTEM_MODULE|CONFIG_MODULE)
    for _, v := range modules {
        if self := v.Self(); self != nil {
            if v.Init(o) == Error {
                return Error
            }
        }
    }

    for _, v := range modules {
        if self := v.Self(); self != nil {
            go v.Main(c)
            //c.Start(v.Main, c)
        }
    }

    return Ok
}

func (r *Module_t) Exit() int {
    if !Sentinel[EXIT] {
        return Ok
    }

    Sentinel[EXIT] = false

    select {

    }

    // TODO Reload

    for _, v := range Modules {
        if self := v.Self(); self != nil {
            v.Exit()
        }
    }

    // TODO close main

    return Ok
}

func (m *Module_t) Self() *Module_t {
    return m
}

func GetModules(m []Module, modType int64) []Module {
    if m == nil || len(m) <= 0 {
        return nil
    }

    pos  := uint(28)
    flag := false

    switch modType {

    case SYSTEM_MODULE:
        pos = 0
        goto PARSE

    case CONFIG_MODULE:
        pos = 0
        goto PARSE

    case SYSTEM_MODULE|CONFIG_MODULE:
        pos = 0; flag = true
        goto PARSE
    }

PARSE:
    var modules []Module

    modType = modType >> pos
    for _, v := range m {
        if self := v.Self(); self != nil {
            moduleType := self.Type >> pos

            if flag {
                if moduleType == SYSTEM_MODULE ||
                   moduleType == CONFIG_MODULE {
                    continue
                }

                modules = append(modules, v)
            } else {
                if moduleType == modType {
                    modules = append(modules, v)
                }
            }
        }
    }

    return modules
}
