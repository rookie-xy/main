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
    CHANNEL_MODULE = 0x20000000
    CHANNEL_CONFIG = MAIN_CONFIG|CONFIG_BLOCK
)

type Channels struct {
    *Channel
    *Module
}

type ChannelsCtx struct {
    *Context
}

var channel = String{ len("channel"), "channel" }
var channelContext = &ChannelsCtx{
    Context: &Context{
        Name: channel,
    },
}

func (cc *ChannelsCtx) Create() unsafe.Pointer {
    return nil
}

func (cc *ChannelsCtx) Insert(p *unsafe.Pointer) int {
    return Ok
}

func (cc *ChannelsCtx) Contexts() *Context {
    return cc.Get()
}

var channels = String{ len("channels"), "channels" }
var channelCommands = []Command{

    { channels,
      CHANNEL_CONFIG,
      channelsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func channelsBlock(cfg *Configure, _ *Command, _ *unsafe.Pointer) int {
    if nil == cfg {
        cfg.Error("channels block error")
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    Block(cfg, Modables, CHANNEL_MODULE, flag)

    return Ok
}

var channelsModule = &Channels{
    Module: &Module{
				    MODULE_V1,
				    CONTEXT_V1,
        channelContext,
				    channelCommands,
        CONFIG_MODULE,
				},

    Channel: NewChannel(),
}

func (c *Channels) Init(o *Option) int {
    fmt.Println("channels init")
    return Ok
}

func (c *Channels) Main(cfg *Configure) int {
    fmt.Println("channels main")
    return Ok
}

func (c *Channels) Exit() int {
    fmt.Println("channels exit")
    return Ok
}

func (c *Channels) Type() *Module {
    return c.Self()
}

func init() {
    Modables = Load(Modables, channelsModule)
}