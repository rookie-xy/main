/*
 * Copyright (C) 2016 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
"fmt"
)

const (
    OUTPUT_MODULE = 0x40000000
    OUTPUT_CONFIG = MAIN_CONFIG|CONFIG_BLOCK
)

type Outputs struct {
    *Channels
}

type Outputable interface {

}

type OutputsCtx struct {
    *Context
}

var output = String{ len("output"), "output" }
var outputContext = &OutputsCtx{
    Context: &Context{
        output,
        nil,
    },
}

func (oc *OutputsCtx) Create() int {
    return Ok
}

func (oc *OutputsCtx) Insert() int {
    return Ok
}

func (oc *OutputsCtx) Type() *Context {
    return oc.Context
}

var outputs = String{ len("outputs"), "outputs" }
var outputCommands = []Command{

    { outputs,
      OUTPUT_CONFIG,
      outputsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func outputsBlock(cfg *Configure, _ *Command, _ *unsafe.Pointer) int {
    /*
    if nil == cycle {
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    cycle.Block(cycle, OUTPUT_MODULE, flag)
    */

    return Ok
}

var outputsModule = &Outputs{
    Channels: &Channels{
        Module: &Module{
            MODULE_V1,
            CONTEXT_V1,
            //unsafe.Pointer(outputContext),
								    outputContext,
            outputCommands,
            CONFIG_MODULE,
        },
    },
}

func (o *Outputs) Init(c *Configure) int {
    fmt.Println("outputs init")
    return Ok
}

func (o *Outputs) Main(c *Channel) int {
    fmt.Println("outputs main")
    return Ok
}

func (o *Outputs) Exit() int {
    fmt.Println("outputs exit")
    return Ok
}

func (o *Outputs) Type() *Module {
    return o.Module
}

func init() {
    Modules = Load(Modules, outputsModule)
}