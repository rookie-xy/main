package types

import (
    "os"
    "unsafe"
)

type Module struct {
    Index      int
    CtxIndex   int
    Context    Context
    Commands   []Command
    Type       int64
}

var Modulers []Moduler

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

func (r *Module) Init(o *Option) int {
    if !Sentinel[INIT] {
        return Ok
    }

    Sentinel[INIT] = false

    modulers := GetSomeModules(Modulers, SYSTEM_MODULE)
    if modulers == nil {
        return Error
    }

    for _, v := range modulers {
        if v != nil {
            if self := v.Self(); self != nil {
                if v.Init(o) == Error {
                    os.Exit(0)
                }
            }
        }
    }

    var configure *Configure
    if configure = o.Configure; configure == nil {
        configure = NewConfigure(o.Log)
    }

    configure.Pointer = unsafe.Pointer(o)

    for _, v := range modulers {
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

    if Block(configure, Modulers, r.Type, CONFIG_BLOCK) == Error {
        return Error
    }

    return Ok
}

func (r *Module) Main(c *Configure) int {
    if !Sentinel[MAIN] {
        return Ok
    }

    Sentinel[MAIN] = false

    o := (*Option)(unsafe.Pointer(uintptr(c.Pointer)))
    if o == nil {
        // TODO add log
        return Error
    }

    modules := GetSpacModules(Modulers)
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
            }
        }
    }

    return Ok
}

func (r *Module) Exit() int {
    if !Sentinel[EXIT] {
        return Ok
    }

    Sentinel[EXIT] = false

    for _, v := range Modulers {
        if v != nil {
            if self := v.Self(); self != nil {
                v.Exit()
            }
        }
    }

    return Ok
}

func (m *Module) Self() *Module {
    return m
}

func Load(modulers []Moduler, moduler Moduler) []Moduler {
    if modulers == nil && moduler == nil {
        return nil
    }

    modulers = append(modulers, moduler)

    return modulers
}

func (m *Module) SetIndex(i int) {
    m.Index = i
}

func GetSomeModules(m []Moduler, modType int64) []Moduler {
    var modulers []Moduler

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

func GetSpacModules(m []Moduler) []Moduler {
    var modulers []Moduler

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

func GetPartModules(m []Moduler, modType int64) []Moduler {
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

    var modulers []Moduler

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

