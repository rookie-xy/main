package main

import (
    . "github.com/rookie-xy/main/types"

    _ "github.com/rookie-xy/main/modules"
    "fmt"
)

var modules = String{ len(Modules), "modules" }

func main() {
    log := NewLog()

    if modules.Len < 1 {
        log.Warn("have not found:", modules.Data)
        fmt.Println("have not found:", modules.Data)
        return
    }

    Modules = Load(Modules, nil)
    for i := 0; Modules[i] != nil; i++ {
        module := Modules[i]
        module.Type().SetIndex(uint(i))
        //Modules[count].Index = uint(count)
    }

    configure := NewConfigure(log)

    Init(Modules, configure)

    channel := NewChannel()

    Main(Modules, channel)

/*    Monitor() */

    Exit(Modules)
}