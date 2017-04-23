/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
)

type Context struct {
    Name   String
    P     *unsafe.Pointer
}

type Contextable interface {
    Create() int
    Insert(p *unsafe.Pointer) int
    Self() *Context
}

func NewContext() *Context {
    return &Context{
        "context",
        nil,
    }
}

func (c *Context) Block(cfg *Configure, modules []Moduleable, modType int64, cfgType int) int {
    for m := 0; modules[m] != nil; m++ {
        module := modules[m].Type()

        if module.Type != modType {
            continue
        }

        if handle := module.Context; handle != nil {
            if this := handle.Create(); this != Error {
                handle.Type().P[module.CtxIndex] = this

            } else {
                return this
            }

        } else {
            continue
        }
    }

    configure := cfg.GetConfigure()
    if configure == nil {
        return Error
    }

    if cfg.SetModuleType(modType) == Error {
        return Error
    }

    if cfg.SetCommandType(cfgType) == Error {
        return Error
    }

    if cfg.Materialized(modules) == Error {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m].Type()
        if module.Type != modType {
            continue
        }

        if handle := module.Context; handle != nil {
            context := handle.Type().P[module.CtxIndex]

            if this := handle.Insert(context); this == Error {
												    return this
            }

        } else {
            continue
        }
    }

    return Ok
}
