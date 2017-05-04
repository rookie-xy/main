package main

import (
    . "github.com/rookie-xy/main/types"

    _ "github.com/rookie-xy/main/modules"

    _ "github.com/rookie-xy/plugins/option/simple/src"
    _ "github.com/rookie-xy/plugins/configure/file/src"

    //_ "github.com/rookie-xy/plugins/inputs/stdin/src"
    _ "github.com/rookie-xy/plugins/inputs/file/src"
    _ "github.com/rookie-xy/plugins/channels/memory/src"
    _ "github.com/rookie-xy/plugins/outputs/stdout/src"

    _ "github.com/rookie-xy/plugins/codecs/plain/src"
    _ "github.com/rookie-xy/plugins/codecs/multiline/src"
    _ "github.com/rookie-xy/plugins/codecs/yaml/src"

    "fmt"
    "os"
)

var modulers = String{ len(Modulers), "modulers" }

func Init(m []Moduler, o *Option) int {
    modulers := GetSomeModules(m, SYSTEM_MODULE)
    if modulers == nil {
        return Error
    }

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

func Main(m []Moduler, o *Option) {
    modules := GetSpacModules(m)

    for i := 0; modules[i] != nil; i++ {
        module := modules[i]
        module.Init(o)
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
        if moduler := Modulers[i]; moduler != nil {
            if module := moduler.Type(); moduler != nil {
                module.SetIndex(i)
            }
        }
    }

    option := NewOption(log)
    if option.SetArgs(len(os.Args), os.Args) == Error {
        return
    }

    Init(Modulers, option)

    Main(Modulers, option)

    Monitor()

    //Exit(Modables)
}