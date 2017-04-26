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
/*
type InputsCtx struct {
    *Context
}

var input = String{ len("input"), "input" }
var inputContext = &InputsCtx{
    Context: &Context{
        Name: input,
    },
}

func (ic *InputsCtx) Create() unsafe.Pointer {
    return nil
}

func (ic *InputsCtx) Insert(p *unsafe.Pointer) int {
    return Ok
}

func (ic *InputsCtx) Contexts() *Context {
    return ic.Get()
}
*/

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
    if nil == cfg {
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    Block(cfg, Modulers, INPUT_MODULE, flag)

    return Ok
}

var inputsModule = &Inputs{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
				    //inputContext,
        nil,
        inputCommands,
        CONFIG_MODULE,
    },
}

func (i *Inputs) Init(o *Option) int {
    fmt.Println("inputs init")
    return Ok
}

func (i *Inputs) Main(c *Configure) int {
    fmt.Println("inputs main")
    return Ok
}

func (i *Inputs) Exit() int {
    fmt.Println("inputs exit")
    return Ok
}

func (i *Inputs) Type() *Module {
    return i.Self()
}

func init() {
    Modulers = Load(Modulers, inputsModule)
}