/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "os"
)

type Module_t struct {
    Index      int
    CtxIndex   int
    Context    Context
    Commands   []Command_t
    Type       int64
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

    modules := GetSomeModules(Modules, SYSTEM_MODULE)
    if modules == nil {
        return Error
    }

    for _, v := range modules {
        if v != nil {
            if self := v.Self(); self != nil {
                if v.Init(o) == Error {
                    os.Exit(0)
                }
            }
        }
    }

    var configure *Configure_t
    if configure = o.Configure_t; configure == nil {
        configure = NewConfigure(o.Log_t)
    }

    for _, v := range modules {
        if v != nil {
            if self := v.Self(); self != nil {
                go v.Main(configure)
            }
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

    modules := GetSpacModules(Modules)
    for _, v := range modules {
        if v != nil {
            if self := v.Self(); self != nil {
                if v.Init(o) == Error {
                    return Error
                }
            }
        }
    }

    for _, v := range modules {
        if v != nil {
            if self := v.Self(); self != nil {
                go v.Main(c)
                //c.Start(v.Main, c)
            }
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
        if v != nil {
            if self := v.Self(); self != nil {
                v.Exit()
            }
        }
    }

    // TODO close main

    return Ok
}

func (m *Module_t) Self() *Module_t {
    return m
}

func Load(modulers []Module, moduler Module) []Module {
    if modulers == nil && moduler == nil {
        return nil
    }

    modulers = append(modulers, moduler)

    return modulers
}

func (m *Module_t) SetIndex(i int) {
    m.Index = i
}

func GetSomeModules(m []Module, modType int64) []Module {
    var modulers []Module

    for _, v := range m {
        if v != nil {
            if self := v.Self(); self != nil {
                if self.Type == modType {
                    modulers = Load(modulers, v)
                }
            }
        }
    }

    modulers = Load(modulers, nil)

    return modulers
}

func GetSpacModules(m []Module) []Module {
    var modulers []Module

    for _, v := range m {
        if v != nil {
            if self := v.Self(); self != nil {
                if self.Type == SYSTEM_MODULE ||
                   self.Type == CONFIG_MODULE {
                    continue
                }

                modulers = Load(modulers, v)
            }
        }
    }

    modulers = Load(modulers, nil)

    return modulers
}

func GetPartModules(m []Module, modType int64) []Module {
    if m == nil || len(m) <= 0 {
        return nil
    }

    switch modType {

    case SYSTEM_MODULE:
        modulers := GetSomeModules(m, modType)
        if modulers != nil {
            return modulers
        }

    case CONFIG_MODULE:
        modulers := GetSomeModules(m, modType)
        if modulers != nil {
            return modulers
        }
    }

    var modulers []Module
    modType = modType >> 28
    for _, v := range m {
        if v != nil {
            if self := v.Self(); self != nil {
                moduleType := self.Type >> 28
                if moduleType == modType {
                    modulers = Load(modulers, v)
                }
            }
        }
    }

    modulers = Load(modulers, nil)

    return modulers
}
