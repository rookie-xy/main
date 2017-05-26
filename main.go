/*
 * Copyright (C) 2017 Meng Shi
 */

package main

import (
      "fmt"
      "os"
    . "github.com/rookie-xy/worker/types"

    _ "github.com/rookie-xy/modules/option/simple/src"
    _ "github.com/rookie-xy/modules/configure/file/src"

    _ "github.com/rookie-xy/modules/inputs/file/src"
    _ "github.com/rookie-xy/modules/channels/topic/src"
    _ "github.com/rookie-xy/modules/outputs/stdout/src"
    _ "github.com/rookie-xy/modules/outputs/elasticsearch/src"

    _ "github.com/rookie-xy/plugins/codecs/plain/src"
    _ "github.com/rookie-xy/plugins/codecs/multiline/src"
    _ "github.com/rookie-xy/plugins/codecs/yaml/src"
)

var (
    modules  = String_t{ len(Modules),    "modules"  }

    channels = String_t{ len("channels"), "channels" }
    inputs   = String_t{ len("inputs"),   "inputs"   }
    outputs  = String_t{ len("outputs"),  "outputs"  }
)

var workerCommands = []Command_t{

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

var worker = &Module_t{
    MODULE_V1,
    CONTEXT_V1,
    nil,
    workerCommands,
    CONFIG_MODULE,
}

func init() {
    Modules = append(Modules, worker)
}

func main() {

    log := NewLog()

    if modules.Len < 1 {
        fmt.Printf("[%s]have not found\n", modules.Data)
        exit()
    }

    for i, v := range Modules {
        if self := v.Self(); self != nil {
            self.SetIndex(i)
        }
    }

    option := NewOption(log)
    if option.SetArgs(len(os.Args), os.Args) == Error {
        exit()
    }

    Sentinel[INIT] = true

    worker.Init(option)

    Sentinel[MAIN] = true

    worker.Main(option.Configure_t)

    Sentinel[EXIT] = true

    worker.Exit()
}

func exit() {
    os.Exit(1)
}
