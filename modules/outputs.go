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
    *Module
}

type OutputsCtx struct {
    *Context
}

var output = String{ len("output"), "output" }
var outputContext = &OutputsCtx{
    Context: &Context{
        Name: output,
    },
}

func (oc *OutputsCtx) Create() unsafe.Pointer {
    return nil
}

func (oc *OutputsCtx) Insert(p *unsafe.Pointer) int {
    return Ok
}

func (oc *OutputsCtx) Contexts() *Context {
    return oc.Get()
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
    if nil == cfg {
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    Block(cfg, Modulers, OUTPUT_MODULE, flag)

    return Ok
}

var outputsModule = &Outputs{
        Module: &Module{
            MODULE_V1,
            CONTEXT_V1,
								    outputContext,
            outputCommands,
            CONFIG_MODULE,
        },
}

func (o *Outputs) Init(opt *Option) int {
    fmt.Println("outputs init")
    return Ok
}

func (o *Outputs) Main(c *Configure) int {
    fmt.Println("outputs main")
    return Ok
}

func (o *Outputs) Exit() int {
    fmt.Println("outputs exit")
    return Ok
}

func (o *Outputs) Type() *Module {
    return o.Self()
}

func init() {
    Modulers = Load(Modulers, outputsModule)
}