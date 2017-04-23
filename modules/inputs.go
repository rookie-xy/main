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
    *Channels
}

type Inputable interface {

}

type InputsCtx struct {
    *Context
}

var input = String{ len("input"), "input" }
var inputContext = &InputsCtx{
    Context: &Context{
        input,
        nil,
    },
}

func (ic *InputsCtx) Create() int {
    return Ok
}

func (ic *InputsCtx) Insert() int {
    return Ok
}

func (ic *InputsCtx) Type() *Context {
    return ic.Context
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

var inputsModule = &Inputs{
    Channels: &Channels{
        Module: &Module{
            MODULE_V1,
            CONTEXT_V1,
            //unsafe.Pointer(inputContext),
								    inputContext,
            inputCommands,
            CONFIG_MODULE,
        },
    },
}

func (i *Inputs) Init(c *Configure) int {
    fmt.Println("inputs init")
    return Ok
}

func (i *Inputs) Main(c *Channel) int {
    fmt.Println("inputs main")
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
    Modules = Load(Modules, inputsModule)
}