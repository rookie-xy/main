package main

import (
    . "github.com/rookie-xy/main/types"

    _ "github.com/rookie-xy/main/modules"

    _ "github.com/rookie-xy/plugins/option/simple/src"
    _ "github.com/rookie-xy/plugins/configure/file/src"
    _ "github.com/rookie-xy/plugins/configure/yaml/src"

    _ "github.com/rookie-xy/plugins/inputs/stdin/src"

    "fmt"
    "os"
)

var modulers = String{ len(Modulers), "modulers" }

func Init(m []Moduler, o *Option) int {
    modulers := GetSomeModules(m, SYSTEM_MODULE)
    if modulers == nil {
        return Error
    }

    fmt.Println(len(modulers))

    for i := 0; modulers[i] != nil; i++ {
        if module := modulers[i]; module != nil {
            if module.Init(o) == Error {
                os.Exit(SYSTEM_MODULE)
            }
        }
    }

    var configure *Configure
    if configure = o.Configure; configure == nil {
        configure = NewConfigure(o.Log)
    }

    for i := 0; modulers[i] != nil; i++ {
        if module := modulers[i]; module != nil {
            go module.Main(configure)
        }
    }

    select {

    case e := <- configure.Event:
        if op := e.GetOpcode(); op != LOAD {
            return Ignore
        }
    }

    if Block(configure, m, CONFIG_MODULE, CONFIG_BLOCK) == Error {
        return Error
    }

    return Ok
}

func Main(m []Moduler, c *Configure) {
    for i := 0; m[i] != nil; i++ {
        module := m[i]
        module.Main(c)
    }
}

func Exit(m []Moduler) {
    for i := 0; m[i] != nil; i++ {
        module := m[i]
        module.Exit()
    }
}

func Monitor() {
    select {

    }
}

func main() {
    log := NewLog()

    if modulers.Len < 1 {
        //log.Warn("have not found:", modables.Data)
        fmt.Printf("[%s]have not found\n", modulers.Data)
        return
    }

    Modulers = Load(Modulers, nil)
    for i := 0; Modulers[i] != nil; i++ {
        if modable := Modulers[i]; modable != nil {
            if module := modable.Type(); modable != nil {
                module.SetIndex(i)
            }
        }
    }

    option := NewOption(log)
    if option.SetArgs(len(os.Args), os.Args) == Error {
        return
    }

    Init(Modulers, option)

    //configure := NewConfigure(log)

    //Main(Modables, configure)

    Monitor()

    //Exit(Modables)
}