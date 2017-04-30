/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
    . "github.com/rookie-xy/main/types"
"fmt"
)

const (
    FILTER_MODULE = 0x60000000
    FILTER_CONFIG = MAIN_CONFIG|CONFIG_BLOCK
)

type Filters struct {
    *Module
}

func (r *Filters) Init(o *Option) int {
    fmt.Println("filters init")
    return Ok
}

func (r *Filters) Main(c *Configure) int {
    fmt.Println("filters main")
    return Ok
}

func (r *Filters) Exit() int {
    fmt.Println("filters exit")
    return Ok
}

func (r *Filters) Type() *Module {
    return r.Self()
}

var filters = String{ len("filters"), "filters" }
var filterCommands = []Command{

    { filters,
      FILTER_CONFIG,
      filterBlock,
      0,
      0,
      nil },

    NilCommand,
}

func filterBlock(c *Configure, _ *Command, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := FILTER_CONFIG|CONFIG_ARRAY
    Block(c, Modulers, FILTER_MODULE, flag)

    return Ok
}

var filterModule = &Filters{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        nil,
        filterCommands,
        CONFIG_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, filterModule)
}