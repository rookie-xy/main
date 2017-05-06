package types

import (
    "fmt"
    "os"
)

type Module struct {
    Index      int
    CtxIndex   int
    Context    Context
    Commands   []Command
    Type       int64
}

var Modulers []Moduler

func (r *Module) Init(o *Option) int {
    modulers := GetSomeModules(Modulers, SYSTEM_MODULE)
    if modulers == nil {
        return Error
    }
/*
    for i := 0; modulers[i] != nil; i++ {
        if module := modulers[i]; module != nil {
            if module.Init(o) == Error {
                os.Exit(SYSTEM_MODULE)
            }
        }
    }
    */
    for _, v := range modulers {
        if v != nil {
            if self := v.Self(); self != nil {
                if v.Init(o) == Error {
                    os.Exit(SYSTEM_MODULE)
                }
            }
        }
    }

    var configure *Configure
    if configure = o.Configure; configure == nil {
        configure = NewConfigure(o.Log)
    }

    /*
    for i := 0; modulers[i] != nil; i++ {
        if module := modulers[i]; module != nil {
            go module.Main(configure)
        }
    }
    */

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

    if Block(configure, Modulers, CONFIG_MODULE, CONFIG_BLOCK) == Error {
        return Error
    }

    return Ok
}

func (r *Module) Main(cfg *Configure) int {
    /*
    modules := GetSpacModules(Modulers)

    for i := 0; modules[i] != nil; i++ {
        module := modules[i]
        module.Init(cfg.Option)
    }
    */

    return Ok
}

func (r *Module) Exit() int {
    fmt.Println("channels exit")
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
/*
    for i := 0; m[i] != nil; i++ {
        module := m[i].Self()

        if module.Type == modType {
            modulers = Load(modulers, m[i])
        }
    }
    */

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
/*
    for i := 0; m[i] != nil; i++ {
        module := m[i].Self()

        if module.Type == SYSTEM_MODULE ||
           module.Type == CONFIG_MODULE {
            continue
        }

        modulers = Load(modulers, m[i])
    }
    */

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

    /*
    for i := 0; m[i] != nil; i++ {
        module := m[i].Self()
        moduleType := module.Type >> 28

        if moduleType == modType {
            modulers = Load(modulers, m[i])
        }
    }
    */

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
