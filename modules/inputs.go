package modules


import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
"fmt"
)

const (
    INPUT_MODULE = 0x30000000
    INPUT_CONFIG = MAIN_CONFIG|CONFIG_BLOCK
)

type Inputs struct {
    *Module
}

var input = String{ len("input"), "input" }
var inputContext = &Context{
    input,
    nil,
    nil,
}

var inputs = String{ len("inputs"), "inputs" }
var inputCommands = []Command{

    { inputs,
      INPUT_CONFIG,
      inputsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func inputsBlock(cfg *Configure, _ *Command, _ *unsafe.Pointer) int {
/*
    if nil == cycle {
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    cycle.Block(cycle, INPUT_MODULE, flag)
    */

    return Ok
}

var inputModule = &Inputs{
    Module: &Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(inputContext),
    inputCommands,
    CONFIG_MODULE,
    },
}

func (i *Inputs) Init() int {
    fmt.Println("inputs init")
    return Ok
}

func (i *Inputs) Main() int {
    fmt.Println("inputs init")
    return Ok
}

func (i *Inputs) Exit() int {
    fmt.Println("inputs exit")
    return Ok
}

func (i *Inputs) Type() *Module {
    return i.Module
}

func init() {
    Modules = Load(Modules, inputModule)
}