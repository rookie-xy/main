package main

import (
    . "github.com/rookie-xy/main/types"

    _ "github.com/rookie-xy/main/modules"
    "fmt"
)

var modules = String{ len(Modules), "modules" }

func main() {
    if modules.Len < 1 {
        fmt.Println("have not found:", modules.Data)
        return
    }

    Modules = Load(Modules, nil)
    for i := 0; Modules[i] != nil; i++ {
        module := Modules[i]
        module.Type().SetIndex(uint(i))
        //Modules[count].Index = uint(count)
    }

    Init(Modules)
/*
    Main(Modules)

    Monitor()

    Exit(Modules)
    */
}