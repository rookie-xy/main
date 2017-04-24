/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
)

type Context struct {
    Name   String
    Data   [1024]*unsafe.Pointer
}

type Contextable interface {
    Create() unsafe.Pointer
    Insert(p *unsafe.Pointer) int

    Contexts() *Context
}

func NewContext() *Context {
    return &Context{
        Name: String{ len("context"), "context" },
    }
}

func (r *Context) Get() *Context {
    return r
}

func (r *Context) Set(c *Context) int {
    /*
    if c == nil {
        return Error
    }

    r = c
    */
    return Ok
}

func (c *Context) SetData(index int, p *unsafe.Pointer) int {
    c.Data[index] = p
    return Ok
}

func (c *Context) GetData(index int) *unsafe.Pointer {
    return c.Data[index]
}

func Block(cfg *Configure, modables []Moduleable, modType int64, cfgType int) int {
    for m := 0; modables[m] != nil; m++ {
        module := modables[m].Type()

        if module.Type != modType {
            continue
        }

        if handle := module.Context; handle != nil {
            if this := handle.Create(); this != nil {
                if context := handle.Contexts(); context != nil {
                    if context.SetData(module.CtxIndex, &this) == Error {
                        return Error
                    }
                }

            } else {
                return Error
            }

        } else {
            continue
        }
    }

    if cfg == nil {
        return Error
    }

    if cfg.SetModuleType(modType) == Error {
        return Error
    }

    if cfg.SetCommandType(cfgType) == Error {
        return Error
    }

    if cfg.Materialized(modables) == Error {
        return Error
    }

    for m := 0; modables[m] != nil; m++ {
        module := modables[m].Type()
        if module.Type != modType {
            continue
        }

        if handle := module.Context; handle != nil {
            if context := handle.Contexts(); context != nil {
                if this := context.GetData(module.CtxIndex); this != nil {
                    if handle.Insert(this) == Error {
                        return Error
                    }
                }
            }

        } else {
            continue
        }
    }

    return Ok
}
