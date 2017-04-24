package main

import (
    . "github.com/rookie-xy/main/types"

    _ "github.com/rookie-xy/main/modules"

    _ "github.com/rookie-xy/module/option/simple/src"
    _ "github.com/rookie-xy/module/configure/file/src"
    _ "github.com/rookie-xy/module/configure/yaml/src"

    "fmt"
    "os"
)

var modables = String{ len(Modables), "modables" }

func Init(modables []Moduleable, o *Option) int {
    modules := GetSomeModules(modables, SYSTEM_MODULE)
    if modules == nil {
        return Error
    }

    for i := 0; modules[i] != nil; i++ {
        if module := modules[i]; module != nil {
            module.Init(o)
        }
    }

    c := NewConfigure(o.Log)

    for i := 0; modules[i] != nil; i++ {
        if module := modules[i]; module != nil {
            module.Main(c)
        }
    }

    select {

    case e := <- c.Event:
        if op := e.GetOpcode(); op != LOAD {
            return Ignore
        }
    }

    if Block(c, modables, CONFIG_MODULE, CONFIG_BLOCK) == Error {
        return Error
    }

    return Ok
}

func Main(modules []Moduleable, c *Configure) {
    for i := 0; modules[i] != nil; i++ {
        module := modules[i]
        module.Main(c)
    }
}

func Exit(modules []Moduleable) {
    for i := 0; modules[i] != nil; i++ {
        module := modules[i]
        module.Exit()
    }
}

func Monitor() {

}

func main() {
    log := NewLog()

    if modables.Len < 1 {
        //log.Warn("have not found:", modables.Data)
        fmt.Printf("[%s]have not found\n", modables.Data)
        return
    }

    Modables = Load(Modables, nil)
    for i := 0; Modables[i] != nil; i++ {
        if modable := Modables[i]; modable != nil {
            if module := modable.Type(); modable != nil {
                module.SetIndex(i)
            }
        }
    }

    option := NewOption(log)
    if option.SetArgs(len(os.Args), os.Args) == Error {
        return
    }

    Init(Modables, option)

    //configure := NewConfigure(log)

    //Main(Modables, configure)

    Monitor()

    //Exit(Modables)
}