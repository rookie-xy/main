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
    CODEC_MODULE = 0x70000000
    CODEC_CONFIG = MAIN_CONFIG|CONFIG_BLOCK
)

type Codecs struct {
    *Module
}

func (r *Codecs) Init(o *Option) int {
    fmt.Println("codecs init")

    return Ok
}

func (r *Codecs) Main(c *Configure) int {
    fmt.Println("codecs main")
    return Ok
}

func (r *Codecs) Exit() int {
    fmt.Println("codecs exit")
    return Ok
}

func (r *Codecs) Type() *Module {
    return r.Self()
}

var codecs = String{ len("codecs"), "codecs" }

var codecCommands = []Command{

    { codecs,
      CODEC_CONFIG,
      codecBlock,
      0,
      0,
      nil },

    NilCommand,
}

func codecBlock(c *Configure, _ *Command, _ *unsafe.Pointer) int {
    if c == nil {
        return Error
    }

    flag := CODEC_CONFIG|CONFIG_MAP
    Block(c, Modulers, CODEC_MODULE, flag)

    return Ok
}

var codecModule = &Codecs{
    Module: &Module{
        MODULE_V1,
        CONTEXT_V1,
        nil,
        codecCommands,
        CONFIG_MODULE,
    },
}

func init() {
    Modulers = Load(Modulers, codecModule)
}
