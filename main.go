package main

import (
    . "github.com/rookie-xy/worker/types"

    _ "github.com/rookie-xy/modules/option/simple/src"
    _ "github.com/rookie-xy/modules/configure/file/src"

    //_ "github.com/rookie-xy/plugins/inputs/stdin/src"
    _ "github.com/rookie-xy/modules/inputs/file/src"
    _ "github.com/rookie-xy/modules/channels/memory/src"
    _ "github.com/rookie-xy/modules/outputs/stdout/src"

    _ "github.com/rookie-xy/plugins/codecs/plain/src"
    _ "github.com/rookie-xy/plugins/codecs/multiline/src"
    _ "github.com/rookie-xy/plugins/codecs/yaml/src"

    "fmt"
    "os"
)

var modulers = String{ len(Modulers), "modulers" }

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

var (
    channels = String{ len("channels"), "channels" }
    inputs   = String{ len("inputs"), "inputs" }
    outputs  = String{ len("outputs"), "outputs" }
)

var programCommands = []Command{

    { channels,
      CHANNEL_CONFIG,
      SetChannel,
      0,
      0,
      nil },


    { inputs,
      INPUT_CONFIG,
      SetInput,
      0,
      0,
      nil },

    { outputs,
      OUTPUT_CONFIG,
      SetOutput,
      0,
      0,
      nil },

    NilCommand,
}

var program = &Module {
    MODULE_V1,
    CONTEXT_V1,
    nil,
    programCommands,
    CONFIG_MODULE,
}

func init() {
    Modulers = Load(Modulers, program)
}

func main() {
    log := NewLog()

    if modulers.Len < 1 {
        fmt.Printf("[%s]have not found\n", modulers.Data)
        return
    }

    Modulers = Load(Modulers, nil)
    for i := 0; Modulers[i] != nil; i++ {
        if moduler := Modulers[i]; moduler != nil {
            if module := moduler.Self(); moduler != nil {
                module.SetIndex(i)
            }
        }
    }

    option := NewOption(log)
    if option.SetArgs(len(os.Args), os.Args) == Error {
        return
    }

    program.Init(option)

    //Init(Modulers, option)

    Main(Modulers, option)

    Monitor()

    //Exit(Modables)
}